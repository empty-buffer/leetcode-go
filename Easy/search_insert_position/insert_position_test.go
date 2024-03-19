package search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     []int
	target   int
	expected int
}

var testCases = []testCase{
	{
		name:     "strait_ok",
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		target:   6,
		expected: 5,
	}, {
		name:     "target_missing_ok",
		data:     []int{1, 2, 3, 4, 5, 7, 8, 9},
		target:   6,
		expected: 5,
	}, {
		name:     "target_missing_ok",
		data:     []int{1, 3, 5, 6},
		target:   5,
		expected: 2,
	},
}

func Test_SearchInsertPositionV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, searchInsertPositionV1(ts.data, ts.target))
		})
	}
}

func Test_Test_SearchInsertPositionV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, searchInsertPositionV2(ts.data, ts.target))
		})
	}
}
