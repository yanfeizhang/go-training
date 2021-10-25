## waitgroup
借助标准库 sync 里的 Waitgroup，可以实现优雅等待所有子 goroutine 完全结束之后主进程才结束退出，这是一种控制并发的方式，可以实现对多 goroutine 的等待。

```go
var wait sync.WaitGroup

func worker(){
	fmt.Println("Worker is working!")
	wait.Done()
}

func main(){
	wait.Add(1)
	go worker()
	wait.Wait()
}
```

## 互斥锁
互斥锁是完全互斥的，两个协程不能进入同一块加锁的区域

```go
//定义锁
lock sync.Mutex

//加锁
lock.Lock()

//解锁
lock.Unlock()
```

### 读写锁
```go
//定义锁
rwlock sync.RWMutex

//加写锁
rwlock.Lock()

//解写锁
rwlock.Unlock()

//加读锁
rwlock.RLock()

//加读锁
rwlock.RUnLock()
```
