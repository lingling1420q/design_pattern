package main

import "fmt"

func printSlice(x []int){
	fmt.Printf("len=%d slice=%v \n" , len(x) , x)
}


func main() {
	//使用range去求一个slice的和
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
}
