package main

import (
	"context"
	"fmt"
)

type userKey int

const (
	someKey1 userKey = iota + 1
	someKey2
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, someKey1, "some value 1")
	ctx = context.WithValue(ctx, someKey2, "some value 2")

	func(ctx context.Context) {
		val, ok := ctx.Value(someKey1).(string)
		if ok {
			fmt.Println("some key", someKey1, " = ", val)
		}
		val, ok = ctx.Value(someKey2).(string)
		if ok {
			fmt.Println("some key", someKey2, " = ", val)
		}
	}(ctx)
}
