package tandemspeed

import (
	"fmt"
	"slices"
	"sort"
)

func TandemSpeed(red []int, blue []int, fastest bool) int {
	if len(red) != len(blue) {
		return 0
	}

	i := len(red) - 1

	slices.Sort(red)
	slices.Sort(blue)

	maxSpeed := 0
	ri := i
	bi := i

	if fastest {
		for ; i >= 0; i-- {
			if red[ri] > blue[bi] {
				maxSpeed += red[ri]
				ri--
			} else {
				maxSpeed += blue[bi]
				bi--
			}
		}
	} else {
		for j := len(red) - 1; j >= 0; j-- {
			if red[j] > blue[j] {
				maxSpeed += red[j]
			} else {
				maxSpeed += blue[j]
			}
		}
	}

	return maxSpeed
}

// Create Varian with reverse slice
func TandemSpeedV2(red []int, blue []int, fastest bool) int {
	if fastest {
		sort.Slice(red, func(i, j int) bool {
			return red[i] > red[j]
		})

		sort.Slice(blue, func(i, j int) bool {
			return blue[i] > blue[j]
		})

	} else {
		sort.Slice(red, func(i, j int) bool {
			return red[i] < red[j]
		})

		sort.Slice(blue, func(i, j int) bool {
			return blue[i] > blue[j]
		})
	}

	total := 0

	l := len(red) - 1

	for i := range len(red) {
		if red[l-i] > blue[i] {
			total += red[l-i]
		} else {
			total += blue[i]
		}
	}

	return total
}

func TestLoops() {
	nums := []int{5, 3, 1, 2, 5, 4}

	for i := 0; i < len(nums); i++ {
		fmt.Printf("non-range loop: %d\n", nums[i])
	}

	for i := range len(nums) {
		fmt.Printf("range loop: %d\n", nums[i])
	}

}
