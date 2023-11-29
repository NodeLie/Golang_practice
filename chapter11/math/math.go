package math

func Average(xs []float64) float64 {
	total := float64(0)
	if len(xs) == 0 {
		return 0
	}
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Нахождение минималнього значения в срезе дробных чисел
func Min(xs []float64) (min float64) {
	if len(xs) == 0 {
		return -1
	}
	min = xs[0]
	for _, v := range xs {
		if min > v {
			min = v
		}
	}
	return
}

// Нахождение максимального значения в срезе дробных чисел
func Max(xs []float64) (max float64) {
	max = xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return
}
