package main

import (
	"context"
	"fmt"
)

// 购物的函数
type handler func(ctx context.Context)

// 做家务事
type interceptor func(ctx context.Context, h handler, ivk invoker)

// 引进来的帮助我们实现拦截器的函数
type invoker func(ctx context.Context, ceps []interceptor, h handler)


func getInvoker(ctx context.Context, ceps []interceptor, ivk invoker, cur int) invoker {
	if len(ceps) - 1 == cur {
		return ivk
	}

	return func(ctx context.Context, ceps []interceptor, h handler) {
		ceps[cur+1](ctx, h, getInvoker(ctx , ceps, ivk, cur+1))
	}
}


//func getInvoker(ctx context.Context, ceps []interceptor, cur int, ivk invoker) invoker {
//	 if cur == len(ceps) - 1 {
//	 	 return ivk
//	 }
//
//	 return func(ctx context.Context, ceps []interceptor, h handler) {
//	 	ceps[cur+1](ctx, h, getInvoker(ctx, ceps, cur+1, ivk))
//	 }
//}

func main() {

	h := func(ctx context.Context) {
		fmt.Println("go to the supermarket ...")
	}

	ivk := func(ctx context.Context, ceps []interceptor, h handler) {
		h(ctx)
	}
	var ceps []interceptor

	inter1 := func(ctx context.Context, h handler, ivk invoker) {
		fmt.Println("clean the floor ... ")
		ivk(ctx, ceps, h)
	}

	inter2 := func(ctx context.Context, h handler, ivk invoker) {
		fmt.Println("wash the clothes ... ")
		ivk(ctx, ceps, h)
	}

	inter3 := func(ctx context.Context, h handler, ivk invoker) {
		fmt.Println("watch TV ... ")
		ivk(ctx, ceps, h)
	}

	ceps = append(ceps, inter1, inter2, inter3)
	ctx := context.Background()

	ceps[0](ctx, h, getInvoker(ctx, ceps, ivk, 0))
}




