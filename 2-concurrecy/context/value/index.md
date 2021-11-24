
context 的很重要的一个作用是在各个子 routine 之间传递一些共享值，例如用户的 id。

使用 context.WithValue 创建一个值节点，并把它们链接到树形结构中。

```go
func main() {
	ctx := context.WithValue(context.Background(), "userid", "1234567")
}
```

只要能获取到该 ctx 的协程都可以获取到该值。

```go
func Handler(ctx context.Context) {
	fmt.Println(ctx.Value("userid"))
	wg.Done()
}
func main() {
	ctx := context.WithValue(context.Background(), "userid", "1234567")
	wg.Add(1)
	go Handler(ctx)
	wg.Wait()
}
```

值得注意，如果 ctx 最终穿起来是一个树形结构。所以所有该 ctx 派生的 context 上执行 WithValue 都能获取到该值。
详细例子参见