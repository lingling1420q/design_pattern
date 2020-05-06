package main

//import "fmt"
//
//type VendingMachine interface {
//
//	// 插入硬币
//	InsertCoin(int)
//
//	// 选择商品
//	ChooseGoods(int)
//
//	// 出货
//	ShipGoods(int)
//
//	// 结算
//	Settlement(int)
//
//}
//
//type vendingMachine struct {
//
//}
//
//func (v *vendingMachine) InsertCoin(state int) {
//
//	switch state {
//	case WaitingForCoin :
//		fmt.Println("插入硬币成功...")
//		break
//	case HasCoin :
//		fmt.Println("售货机已经有硬币，请等待...")
//		break
//	case WaitingForShipGoods :
//		fmt.Println("商品选择中，请等待购买完成之后再投币...")
//		break
//	case Refund :
//		fmt.Println("结算中，请耐心等待...")
//		break
//	case FreeGoods :
//		fmt.Println("免费水果，无需硬币购买...")
//		break
//	}
//}
//
//func (v *vendingMachine) ChooseGoods(state int) {
//
//	switch state {
//	case WaitingForCoin :
//		fmt.Println("请先插入硬币...")
//		break
//	case HasCoin :
//		fmt.Println("请选择商品...")
//		break
//	case WaitingForShipGoods :
//		fmt.Println("请先选择商品...")
//		break
//	case Refund :
//		fmt.Println("结算中，请耐心等待...")
//		break
//	case FreeGoods :
//		fmt.Println("免费水果，每人只能选择一份哦...")
//		break
//	}
//
//}
//
//func (v *vendingMachine) ShipGoods(state int) {
//
//	switch state {
//	case WaitingForCoin :
//		fmt.Println("请先插入硬币...")
//		break
//	case HasCoin :
//		fmt.Println("请先选择商品...")
//		break
//	case WaitingForShipGoods :
//		fmt.Println("出货中... 请耐心等待")
//		break
//	case Refund :
//		fmt.Println("结算中，请耐心等待结算完成再操作...")
//		break
//	case FreeGoods :
//		fmt.Println("出货中... 请耐心等待")
//		break
//	}
//
//}
//
//func (v *vendingMachine) Settlement(state int) {
//
//	switch state {
//	case WaitingForCoin :
//		fmt.Println("清先插入硬币...")
//		break
//	case HasCoin :
//		fmt.Println("请先选择商品...")
//		break
//	case WaitingForShipGoods :
//		fmt.Println("请耐心等待出货...")
//		break
//	case Refund :
//		fmt.Println("结算完成,请拿走您的商品...")
//		break
//	case FreeGoods :
//		fmt.Println("结算完成,请拿走您的免费水果...")
//		break
//	}
//
//}
//
//type goods struct {
//	state int
//}
//
//const WaitingForCoin = 1
//const HasCoin = 2
//const WaitingForShipGoods = 3
//const Refund = 4
//
//const FreeGoods = 5
//
//func buyGoods() {
//	g := &goods{
//		state : WaitingForCoin,
//	}
//	v := &vendingMachine{}
//	v.InsertCoin(g.state)
//	g.state = HasCoin
//	v.ChooseGoods(g.state)
//	g.state = WaitingForShipGoods
//	v.ShipGoods(g.state)
//	g.state = Refund
//	v.Settlement(g.state)
//}
//
//func main() {
//	buyGoods()
//}