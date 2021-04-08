package main

import (
	"fmt"
	lca "github.kiragoo.com/golang-coding/lc/array"
)

func main() {

	//fmt.Println(lct.PreOrderTraversal(
	//	&lct.Node{
	//		Val: 1,
	//		Left: &lct.Node{
	//			Val: 2,
	//			Left: nil,
	//			Right: nil,
	//		},
	//		Right: &lct.Node{
	//			Val: 3,
	//			Left: nil,
	//			Right: nil,
	//		},
	//	}, ))

	fmt.Println(lca.FindMin([]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}))

}
