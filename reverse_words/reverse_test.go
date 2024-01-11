package reverse_words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseWords(t *testing.T) {
	assert.Equal(t, "light a is the", reverseWords("the is a light"))
}
