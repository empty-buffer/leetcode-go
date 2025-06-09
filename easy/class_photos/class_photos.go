package classphotos

import (
	"fmt"
	"sort"
)

func classPhoto(red, blue []int) bool {
	if len(red) != len(blue) {
		return false
	}

	sort.Ints(red)
	sort.Ints(blue)

	if red[len(red)-1] > blue[len(blue)-1] {
		for i := len(red) - 1; i >= 0; i-- {
			if blue[i] < red[i] {
				continue
			} else {
				return false
			}
		}
	} else {
		for i := len(blue) - 1; i >= 0; i-- {
			fmt.Println(blue[i], red[i])
			if blue[i] > red[i] {
				continue
			} else {
				return false
			}
		}
	}

	return true
}
