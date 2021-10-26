package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//设置2秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	context.TODO()
	defer cancel()

	go handle(ctx)
	time.Sleep(time.Second * 5)
	fmt.Println("main routine end")
}

func handle(ctx context.Context) {
	ch := make(chan struct{}, 0)
	go func() {
		//设置1秒耗时的时候，子进程可以正常结束
		//设置4秒耗时任务的时候，会触发ctx的超时
		time.Sleep(time.Second * 1)
		//time.Sleep(time.Second * 4)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		fmt.Println("sub routine done")
	case <-ctx.Done():
		fmt.Println("sub routine timeout")
	}
}
