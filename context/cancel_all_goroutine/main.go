package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	ctxParent, cancel := context.WithCancel(ctx)
	go parent(ctxParent, "Hello-parent")

	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func parent(ctx context.Context, str string) {
	childCtx, cancel := context.WithCancel(ctx)
	go child(childCtx, "Hello-child")
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), str)
			return
		}
	}
}

func child(ctx context.Context, str string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err(), str)
			return
		}
	}
}
