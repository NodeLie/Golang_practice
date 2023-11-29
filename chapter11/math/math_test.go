package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var testsForAverage = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
	{[]float64{}, 0},
}
var testsForMin = []testpair{
	{[]float64{1, 2}, 1},
	{[]float64{1, 1, 1, 1, 1}, 1},
	{[]float64{1}, 1},
	{[]float64{}, -1},
}

func TestAverage(t *testing.T) {
	var v float64
	for _, pair := range testsForAverage {
		v = Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"excepted", pair.average,
				"got", v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	var v float64
	for _, pair := range testsForMin {
		v = Min(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"excepted", pair.average,
				"got", v,
			)
		}
	}
}
