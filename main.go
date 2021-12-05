package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go context tutorial")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println(ctx.Err())
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
		fmt.Println(ctx.Err())
	}

	time.Sleep(2 * time.Second)
}

func doSomethingCool(ctx context.Context) {
	rId := ctx.Value("request-id")
	fmt.Println(rId)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("Doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}
