package common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     []string
	expected string
}

var testCases = []testCase{
	{
		name:     "FL",
		data:     []string{"flower", "flow", "flight"},
		expected: "fl",
	}, {
		name:     "No common Char",
		data:     []string{"dog", "racecar", "car"},
		expected: "",
	}, {
		name:     "Empty",
		data:     []string{},
		expected: "",
	},
}

func Test_CommonPrefixV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, longestCommonPrefixV1(ts.data))
		})
	}
}

func Test_CommonPrefixV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, longestCommonPrefixV2(ts.data))
		})
	}
}

func Test_CommonPrefixV3(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, longestCommonPrefixV3(ts.data))
		})
	}
}
