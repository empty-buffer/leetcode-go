package optimalfreelancing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OptimalFreelancing(t *testing.T) {
	jobs := []Job{
		{
			time:    1,
			paymant: 13,
		},
		{
			time:    2,
			paymant: 23,
		},
		{
			time:    2,
			paymant: 23,
		},
		{
			time:    14,
			paymant: 13,
		},
		{
			time:    4,
			paymant: 33,
		},
		{
			time:    4,
			paymant: 53,
		},
		{
			time:    3,
			paymant: 13,
		},
	}

	assert.Equal(t, 145, OptimalWeek(jobs))
}
