package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "req-id", "55120")
}

func doSomething(ctx context.Context) {
	rId := ctx.Value("req-id")

	fmt.Println(rId)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something")

		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	//ctx := context.Background()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	//cancel() //means deadline is canclled so , <-ctx.Done() will be true
	//if we excuted above statement than this fmt.Println(ctx.Err()) //it will show "context canceled"

	defer cancel() //means deadline will be excuted by it self till its entire cycle

	fmt.Println(ctx.Err()) // but here it is not excuted so it will return "<nil>" as it yet to excute

	ctx = enrichContext(ctx)

	go doSomething(ctx)

	select {

	case <-ctx.Done():
		fmt.Println("oh no, i have exceeded")
		fmt.Println(ctx.Err()) // no error so it will return " context deadline exceeded "

	}

	time.Sleep(2 * time.Second)

}
