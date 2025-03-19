package main

import "fmt"

func Sum(nums ...int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}
