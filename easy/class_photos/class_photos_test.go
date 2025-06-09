package classphotos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClassPhotos(t *testing.T) {
	type tcase struct {
		name string
		red  []int
		blue []int
		res  bool
	}

	cases := []tcase{
		{
			name: "true",
			red:  []int{5, 8, 1, 3, 4},
			blue: []int{6, 9, 5, 2, 4},
			res:  true,
		},
		{
			name: "false",
			red:  []int{5, 8, 1, 3, 5},
			blue: []int{6, 9, 5, 2, 4},
			res:  false,
		},
		{
			name: "not equal len",
			red:  []int{5, 8, 1, 3, 5, 1},
			blue: []int{6, 9, 5, 2, 4},
			res:  false,
		},
	}

	for _, ts := range cases {
		t.Run(ts.name, func(t *testing.T) {
			assert.Equal(t, ts.res, classPhoto(ts.red, ts.blue))
		})
	}

}
