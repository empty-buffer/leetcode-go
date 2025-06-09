package first_occurrence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     string
	target   string
	expected int
}

var testCases = []testCase{
	{
		name:     "Test_one_ok",
		target:   "sad",
		data:     "sadbutsad",
		expected: 0,
	}, {
		name:     "Test_two_ok",
		target:   "sad",
		data:     "butsad",
		expected: 3,
	}, {
		name:     "Test_three_not_ok",
		target:   "leetcode",
		data:     "leed",
		expected: -1,
	},
}

func Test_(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, firstOccurrenceV1(ts.data, ts.target))
		})
	}
}
