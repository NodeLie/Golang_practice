package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func testZam() {
	x := 0
	y := 1
	add := func() int {
		return x + y
	}
	fmt.Println(add())
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2st")
}
func main() {
	xs := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(xs))
	testZam()
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	defer second()
	first()
}
