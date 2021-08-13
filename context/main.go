package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(c context.Context) {
	select {
	case <-c.Done():
		err := c.Err()
		if err != nil {
			fmt.Println(err)
		}
	case <-time.After(3 * time.Second):
		fmt.Println("Time After")
		fmt.Println("do something done")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("cancel")
		cancel()
	}()
	doSomething(ctx)
}
