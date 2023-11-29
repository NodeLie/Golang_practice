package main

import (
	m "chapter11/math"
	"chapter11/testPack"
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	arr := []int{1, 2, 3}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(testPack.Average10(arr))
	fmt.Println(m.Min([]float64{1, 2}))
}
