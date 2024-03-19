package merge_strings_alternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		w1       string
		w2       string
		expected string
	}{
		{
			name:     "Equal",
			w1:       "123",
			w2:       "456",
			expected: "142536",
		}, {
			name:     "Small w1",
			w1:       "1",
			w2:       "456",
			expected: "1456",
		}, {
			name:     "Small w2",
			w1:       "123",
			w2:       "4",
			expected: "1423",
		}, {
			name:     "Empty w1",
			w1:       "",
			w2:       "456",
			expected: "456",
		}, {
			name:     "Empty w2",
			w1:       "123",
			w2:       "",
			expected: "123",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := MergeStringsAlternately(test.w1, test.w2)
			assert.Equal(t, test.expected, res)
			//fmt.Println(res)
		})
	}
}
