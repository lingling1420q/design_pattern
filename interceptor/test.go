package main
//
//import (
//	"context"
//	"fmt"
//)
//
//type interceptor func(ctx context.Context, h handler, ivk invoker)
//
//type handler func(ctx context.Context)
//
//
//func main() {
//
//	var ctx context.Context
//
//	var ceps []interceptor
//
//	var ivk = func(ctx context.Context, interceptors []interceptor, h handler) {
//		h(ctx)
//	}
//
//	var h = func(ctx context.Context) {
//		fmt.Println("go to the supermarket...")
//	}
//
//	var inter1 = func(ctx context.Context, h handler, ivk invoker) {
//		fmt.Println("clean the floor...")
//		ivk(ctx, ceps, h)
//	}
//
//	var inter2 = func(ctx context.Context, h handler, ivk invoker) {
//		fmt.Println("wash the clothes ...")
//		ivk(ctx, ceps, h)
//	}
//
//	var inter3 = func(ctx context.Context, h handler, ivk invoker) {
//		fmt.Println("watch TV ...")
//		ivk(ctx, ceps, h)
//	}
//
//
//	ceps = append(ceps, inter1, inter2, inter3)
//
//	ceps[0](ctx, h, getInvoker(ctx, ceps,0, ivk))
//}
//
//
//
//
//type invoker func(ctx context.Context, interceptors []interceptor, h handler)
//
//
//func getInvoker(ctx context.Context, interceptors []interceptor, cur int, ivk invoker) invoker {
//	if cur == len(interceptors) - 1{
//		return ivk
//	}
//
//	return func(ctx context.Context, interceptors []interceptor, h handler) {
//		interceptors[cur+1](ctx, h, getInvoker(ctx, interceptors, cur+1, ivk))
//	}
//}