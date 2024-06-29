package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "order-id", "12457")
	ctx = context.WithValue(ctx, "order-type", "sell")
	return ctx
}
func getContextValue(ctx context.Context) {
	var output = make(map[any]any) //map[any]any
	output["order-id"] = (ctx.Value("order-id"))
	output["order-type"] = (ctx.Value("order-type"))
	fmt.Println("output[\"order-id\"]", output["order-id"])
	i := 0
	for {
		i++
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing somethind cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
	// return output
}

func main() {
	fmt.Println("go context tutorial")
	// ctx := context.Background() // option 1
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go getContextValue(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("oh no, I've exceeded the deadline")
		fmt.Println(ctx.Err())
	}

	time.Sleep(2 * time.Second)
	// ctxattr:= getContextValue(ctx)
	// for i, v := range ctxattr {
	// 	fmt.Println("i,v:", i, v)
	// }
}
