package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     string
	expected int
}

var testCases = []testCase{
	{
		name:     "MCMXCIV",
		data:     "MCMXCIV",
		expected: 1994,
	}, {
		name:     "III",
		data:     "III",
		expected: 3,
	},
}

func Test_RomanToIntV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, romanToIntV1(ts.data))
		})
	}
}

func Test_Test_RomanToIntV2(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, romanToIntV2(ts.data))
		})
	}
}
