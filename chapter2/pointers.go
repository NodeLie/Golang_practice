package main

import "fmt"

func zero(xPtr *int) {
	*xPtr += 1
}
func square(x *float64) {
	*x = *x * *x
}
func swap(firstValue, secondValue *int) {
	fmt.Println(*firstValue, *secondValue)
	temp := *firstValue
	*firstValue = *secondValue
	*secondValue = temp
}
func main() {
	x := 1.5
	square(&x)
	//zero(&x)
	fmt.Println(x)
	y := 1
	z := 3
	swap(&y, &z)
	fmt.Println(y)
	fmt.Println(z)
}
