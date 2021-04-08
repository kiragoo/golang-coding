package main

import "fmt"

type processFun func(int) bool

func main() {

	//slice := []int{1,2,3,4,5}
	//fmt.Println("slice= ",slice)
	//odd := filter(slice, isOdd)
	//fmt.Println("odd:= ", odd)
	//even := filter(slice, isEven)
	//fmt.Println("even:= ", even)

	//replaceStr("fuck")
	nums := []int{1, 2, 3, 4}
	for _, v := range nums {
		defer func() { fmt.Println(v) }()
	}

}

func replaceStr(str string) {
	byteS := []byte(str)
	byteS[0] = 'F'
	fmt.Println(string(byteS))

}

func isEven(integer int) bool {
	return integer%2 != 0
}

func isOdd(integer int) bool {
	return integer%2 == 0
}

func filter(nums []int, f processFun) []int {
	var result []int
	for _, v := range nums {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
