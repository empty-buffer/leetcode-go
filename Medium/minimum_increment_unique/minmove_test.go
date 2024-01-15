package minimum_increment_unique

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	data           []int
	expectedResult int
}

var testCases = []testCase{
	{
		name:           "Sequential",
		data:           []int{1, 2, 2, 3, 4, 5},
		expectedResult: 4,
	},
	{
		name:           "Leetcode",
		data:           []int{3, 2, 1, 2, 1, 7},
		expectedResult: 6,
	},
}

func Test_MinMoveV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expectedResult, minMoveV1(ts.data))
		})
	}
}

func Test_MinMoveV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expectedResult, minMoveV2(ts.data))
		})
	}
}
