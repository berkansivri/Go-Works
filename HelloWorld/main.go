package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
	OddEven()
}

// OddEven ...
func OddEven() {
	nums := []int{}
	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}
	for _, n := range nums {
		if n%2 == 0 {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
	}
}
