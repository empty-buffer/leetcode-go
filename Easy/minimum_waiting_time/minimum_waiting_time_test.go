package minimumwaitingtime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumWaitingTime(t *testing.T) {
	nums := []int{3, 2, 1, 2, 6}

	assert.Equal(t, 17, MinimumWaitingTime(nums))
}
