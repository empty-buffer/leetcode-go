package rotate_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	data           []int
	steps          int
	expectedResult []int
}

var VeryLongCase = []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 7, 6, 5}

var testCases = []testCase{
	{
		name:           "Simple",
		data:           []int{1, 2, 3, 4, 5, 6, 7},
		steps:          3,
		expectedResult: []int{5, 6, 7, 1, 2, 3, 4},
	}, {
		name:           "Small",
		data:           []int{1, 2, 3, 4},
		steps:          6,
		expectedResult: []int{3, 4, 1, 2},
	}, {
		name:           "Negative",
		data:           []int{99, -1, -100, 3, 55, -12},
		steps:          2,
		expectedResult: []int{55, -12, 99, -1, -100, 3},
	},
}

func Test_RotateV1(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			res := rotateV1(ts.data, ts.steps)
			fmt.Println(res)
			assert.Equal(t, ts.expectedResult, res)
		})
	}
}

func Test_RotateV3(t *testing.T) {
	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			rotateV3(ts.data, ts.steps)
			fmt.Println(ts.data)
			assert.Equal(t, ts.expectedResult, ts.data)
		})
	}
}

func Test_RotateV4(t *testing.T) {
	s := 14
	s %= 6
	fmt.Println("steps", s)

	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			rotateV4(ts.data, ts.steps)
			fmt.Println(ts.data)
			assert.Equal(t, ts.expectedResult, ts.data)
		})
	}
}

func BenchmarkRotateV1(b *testing.B) {
	rotateV1(VeryLongCase, testCases[0].steps)
}

func BenchmarkRotateV3(b *testing.B) {
	rotateV3(VeryLongCase, testCases[0].steps)
}

// WINNER
func BenchmarkRotateV4(b *testing.B) {
	rotateV4(VeryLongCase, testCases[0].steps)
}
