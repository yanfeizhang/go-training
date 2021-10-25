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