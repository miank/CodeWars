package main

import "fmt"

// If trying to paste the same solution in code wars, remove * from Maps function.
func Maps(x *[]int) []int {
	arr := []int{}
	for i := 0; i < len(*x); i++ {
		arr = append(arr, (*x)[i]*2)
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = Maps(&arr)
	fmt.Println(arr)
}
