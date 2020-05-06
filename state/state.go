package main

import (
	"context"
	"fmt"
)

// 售货机
type VendingMachine interface {
	// 投入硬币
	InsertCoin()
	// 选择商品
	ChooseGoods()
	// 出货
	ShipGoods()
	// 结算
	Settlement()

}

type State interface {
	// 投入硬币
	InsertCoin()
	// 选择商品
	ChooseGoods()
	// 出货
	ShipGoods()
	// 结算
	Settlement()
}

type waitingForCoin struct {

}

func (w *waitingForCoin) InsertCoin() {
	fmt.Println("插入硬币成功...")
}

func (w *waitingForCoin) ChooseGoods() {
	fmt.Println("请先插入硬币...")
}

func (w *waitingForCoin) ShipGoods() {
	fmt.Println("请先插入硬币...")
}

func (w *waitingForCoin) Settlement() {
	fmt.Println("清先插入硬币...")
}

type existsCoin struct {

}

func (w *existsCoin) InsertCoin() {
	fmt.Println("售货机已经有硬币，请等待...")
}

func (w *existsCoin) ChooseGoods() {
	fmt.Println("请选择商品...")
}

func (w *existsCoin) ShipGoods() {
	fmt.Println("请先插入硬币...")
}

func (w *existsCoin) Settlement() {
	fmt.Println("请先选择商品...")
}

type waitingForShipGoods struct {

}

func (w *waitingForShipGoods) InsertCoin() {
	fmt.Println("商品选择中，请等待购买完成之后再投币...")
}

func (w *waitingForShipGoods) ChooseGoods() {
	fmt.Println("请先选择商品...")
}

func (w *waitingForShipGoods) ShipGoods() {
	fmt.Println("出货中... 请耐心等待")
}

func (w *waitingForShipGoods) Settlement() {
	fmt.Println("请耐心等待出货...")
}

type freeGoods struct {

}

// 等待投币
const WaitingForCoin = 1
// 投币完成
const ExistsCoin = 2

const WaitingForShipGoods = 3
// 结算
const Refund = 4


func buyGoods() {
	w := &waitingForCoin{}
	w.InsertCoin()
	e := &existsCoin{}
	e.ChooseGoods()
	c := &waitingForShipGoods{}
	c.ShipGoods()

}

func main() {
	buyGoods()

	context.Background()
}