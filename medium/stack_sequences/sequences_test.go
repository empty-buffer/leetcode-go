package stack_sequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	pushData       []int
	popData        []int
	expectedResult bool
}

var testCases = []testCase{
	{
		name:           "Correct Sequence",
		pushData:       []int{1, 2, 3, 4, 5},
		popData:        []int{4, 5, 3, 2, 1},
		expectedResult: true,
	}, {
		name:           "Incorrect Sequence",
		pushData:       []int{1, 2, 3, 4, 5},
		popData:        []int{4, 5, 3, 1, 2},
		expectedResult: false,
	}, {
		name:           "Small Incorrect Sequence",
		pushData:       []int{2, 1, 0},
		popData:        []int{1, 2, 0},
		expectedResult: true,
	},
}

func Test_SequenceV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expectedResult, sequencesV1(ts.pushData, ts.popData))
		})
	}
}

func Test_SequenceV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expectedResult, sequencesV2(ts.pushData, ts.popData))
		})
	}
}
