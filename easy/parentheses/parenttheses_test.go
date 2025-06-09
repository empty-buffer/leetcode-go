package parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	data     string
	expected bool
}

var testCases = []testCase{
	{
		name:     "Correct Simple",
		data:     "(){}[]",
		expected: true,
	}, {
		name:     "Correct Complex",
		data:     "({[]})",
		expected: true,
	}, {
		name:     "Correct Complex 2",
		data:     "({[]{}[]})",
		expected: true,
	}, {
		name:     "Incorrect",
		data:     "({[{}[]})",
		expected: false,
	},
}

func Test_(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.expected, parenthesesV1(ts.data))
		})
	}
}
