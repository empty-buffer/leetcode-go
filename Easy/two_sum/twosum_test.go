package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     []int
	target   int
	expected []int
}

var testCases = []testCase{
	{
		name:     "Basic",
		data:     []int{2, 7, 11, 15},
		target:   13,
		expected: []int{0, 2},
	}, {
		name:     "Same Number",
		data:     []int{0, 3, 3},
		target:   6,
		expected: []int{1, 2},
	},
}

func Test_TwoSumV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, twoSumV1(ts.data, ts.target))
		})
	}
}

func Test_TwoSumV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, twoSumV2(ts.data, ts.target))
		})
	}
}
