package main

import "fmt"

func sum(srez []int) (result int) {
	for _, v := range srez {
		result += v
	}
	return
}
func half(number int) (int, bool) {
	result := number / 2
	isEven := result%2 == 0
	return result, isEven
}
func manyParams(args ...int) (max int) {
	max = args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return
}
func makeOddGenerator() func() uint {
	odd_number := uint(1)
	return func() (res uint) {
		res = odd_number
		odd_number += 2
		return
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func main() {
	srez := []int{1, 2, 5}
	fmt.Println(sum(srez))

	result1, isEven1 := half(1)
	fmt.Println(result1, isEven1)

	result2, isEven2 := half(2)
	fmt.Println(result2, isEven2)

	fmt.Println(manyParams(1, 2, 3))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fib(6))
}
