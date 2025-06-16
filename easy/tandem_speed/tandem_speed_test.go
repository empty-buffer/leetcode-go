package tandemspeed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TandemSpeed(t *testing.T) {
	testCases := []struct {
		name     string
		maxSpeed bool
	}{
		{
			name:     "fast",
			maxSpeed: true,
		},
		{
			name:     "slow",
			maxSpeed: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			r := []int{5, 5, 3, 9, 2}
			b := []int{3, 6, 7, 1, 2}
			if tt.maxSpeed {
				assert.Equal(t, 32, TandemSpeed(r, b, tt.maxSpeed))
			} else {

				assert.Equal(t, 25, TandemSpeed(r, b, tt.maxSpeed))
			}
		})
	}
}

func Test_TandemSpeedV2(t *testing.T) {
	testCases := []struct {
		name     string
		maxSpeed bool
	}{
		{
			name:     "fast",
			maxSpeed: true,
		},
		{
			name:     "slow",
			maxSpeed: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			r := []int{5, 5, 3, 9, 2}
			b := []int{3, 6, 7, 1, 2}
			if tt.maxSpeed {
				assert.Equal(t, 32, TandemSpeedV2(r, b, tt.maxSpeed))
			} else {

				assert.Equal(t, 25, TandemSpeed(r, b, tt.maxSpeed))
			}
		})
	}
}

func Test_Loops(t *testing.T) {
	TestLoops()
}
