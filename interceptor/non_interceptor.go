package main

//// 购物的函数
//type handler func(ctx context.Context)
//
//func main() {
//
//	h := func(ctx context.Context) {
//		// 去购物
//		fmt.Println("shopping....")
//	}
//
//	// 第一件家务事
//	washFloor := func(ctx context.Context, h handler) {
//		fmt.Println("wash the floor...")
//		h(ctx)
//	}
//
//	// 第二件家务事 : 洗衣服
//	washClothes := func(ctx context.Context, h handler) {
//		fmt.Println("wash the clothes ...")
//		h(ctx)
//	}
//
//	// 第三件事 ： 看电视剧
//	watchTV := func(ctx context.Context, h handler) {
//		fmt.Println("watch TV ...")
//		h(ctx)
//	}
//
//	ctx := context.Background()
//	washFloor(ctx, h)
//	washClothes(ctx, h)
//	watchTV(ctx, h)
//}

