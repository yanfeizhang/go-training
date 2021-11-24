package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func HandlerA(ctx context.Context) {
	fmt.Println(ctx.Value("userid"))
	subctx := context.WithValue(ctx, "userip", "127.0.0.1")
	go HandlerB(subctx)
	wg.Done()
}

func HandlerB(ctx context.Context) {
	fmt.Println(ctx.Value("userid"))
	fmt.Println(ctx.Value("userip"))
	wg.Done()
}

func main() {
	ctx := context.WithValue(context.Background(), "userid", "1234567")
	wg.Add(2)
	go HandlerA(ctx)
	wg.Wait()
}
