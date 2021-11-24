package main

import (
	"fmt"
	"sync"
	"time"
)

// 1.1 接口定义
// Context 接口
type Context interface {
	Deadline()
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

// canceler 接口
type canceler interface {
	cancel()
	Done() <-chan struct{}
}

// 1.2 emptyCtx 类型
type emptyCtx int

func (*emptyCtx) Deadline() {
	fmt.Println("*emptyCtx's Deadline")
	return
}
func (*emptyCtx) Done() <-chan struct{} {
	fmt.Println("*emptyCtx's Done")
	return nil
}
func (*emptyCtx) Err() error {
	fmt.Println("*emptyCtx's Err")
	return nil
}
func (*emptyCtx) Value(key interface{}) interface{} {
	fmt.Println("*emptyCtx's Value")
	return nil
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

//1.3 cancelCtx类型
//注意：它只实现了 Value、Done、Err，并没有实现 Deadline，这个依赖其匿名成员 Context 实现的方法。
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
}

func (c *cancelCtx) Value(key interface{}) interface{} {
	fmt.Println("*cancelCtx's Value")
	return c.Context.Value(key)
}

func (c *cancelCtx) Done() <-chan struct{} {
	fmt.Println("*cancelCtx's Done")
	if c.done == nil {
		c.done = make(chan struct{})
	}
	d := c.done
	return d
}

func (c *cancelCtx) Err() error {
	fmt.Println("*cancelCtx's Err")
	err := c.err
	return err
}

//注意，这是 cancel 接口实现的方法
func (c *cancelCtx) cancel() {
	fmt.Println("*cancelCtx's cancel")
}

//1.4 timerCtx 类型
//注意：它只实现了 Deadline，并没有实现 Value、Done、Err，这些依赖其匿名成员 cancelCtx 实现的方法。
type timerCtx struct {
	cancelCtx
	timer    *time.Timer
	deadline time.Time
}

func (c *timerCtx) Deadline() {
	fmt.Println("*timerCtx's Deadline")
}

//注意，这是 cancel 接口实现的方法
func (c *timerCtx) cancel() {
	fmt.Println("*timerCtx's cancel")
}

//2.1 context.Background 和 context.Todo
func Background() Context {
	return background
}
func TODO() Context {
	return todo
}

//2.2 WithCancel
type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	return &c, func() { c.cancel() }
}
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}

//2.3 WithTimeout 函数
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	//初始化一个 timerCtx 对象出来
	//使用 newCancelCtx 初始化 cancelCtx，deadline 就设置为过期时间
	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	return c, func() { c.cancel() }
}

func main() {
	//1. 查看 Background 和 TODO
	fmt.Println(Background())
	fmt.Println(TODO())

	//2. 查看 cancelCtx
	ctx1, func1 := WithCancel(Background())
	fmt.Println("以下是 cancelCtx 的各个方法的执行过程")
	ctx1.Done()
	ctx1.Deadline()
	ctx1.Err()
	ctx1.Value("anykey")
	defer func1() //会调用 cancelCtx.cancel

	//3. 查看 timerCtx
	ctx2, func2 := WithTimeout(Background(), 86400)
	fmt.Println("以下是 timerCtx 的各个方法的执行过程")
	ctx2.Done()
	ctx2.Deadline()
	ctx2.Err()
	ctx2.Value("anykey")
	defer func2() //会调用 timerCtx.cancel
}
