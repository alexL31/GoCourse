package statistic

import (
	"GoCourse/master/home-work6/test/statistic"
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := statistic.Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
func TestSum(t *testing.T) {
	var v float64
	v = statistic.Sum([]float64{1, 2})
	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}
