package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     int
	expected bool
}

var testCases = []testCase{
	{
		name:     "Odd Number",
		data:     121,
		expected: true,
	}, {
		name:     "Long Number",
		data:     12233221,
		expected: true,
	}, {
		name:     "A negative number",
		data:     -12233221,
		expected: false,
	},
}

func Test_PalindromeV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, palindromeV1(ts.data))
		})
	}
}
