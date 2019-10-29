package math

import (
	"github.com/pierdipi/go/common/assert"
	"testing"
)

func TestMaxInt(t *testing.T) {
	tt := []struct {
		a, b, expected int
	}{
		{a: -1, b: -1, expected: -1},
		{a: -1, b: 0, expected: 0},
		{a: -100, b: -99, expected: -99},
		{a: 0, b: 1, expected: 1},
	}
	assertor := assert.New(t)
	var max int
	for r := range tt {
		max = MaxInt(tt[r].a, tt[r].b)
		assertor.AssertThatInt(max).IsEqualInt(tt[r].expected)
	}
}

func TestMinInt(t *testing.T) {
	tt := []struct {
		a, b, expected int
	}{
		{a: -1, b: -1, expected: -1},
		{a: -1, b: 0, expected: -1},
		{a: -100, b: -99, expected: -100},
		{a: 0, b: 1, expected: 0},
	}
	assertor := assert.New(t)
	var min int
	for r := range tt {
		min = MinInt(tt[r].a, tt[r].b)
		assertor.AssertThatInt(min).IsEqualInt(tt[r].expected)
	}
}
