package depthfirstsearch

import (
	"fmt"
	"testing"
)

func Test_DepthFirstSearch(t *testing.T) {
	Tree := NewNode("a").
		AddChildren(
			NewNode("b").AddChildren(
				NewNode("e").AddChildren(
					NewNode("g").Build(),
				).Build(),
				NewNode("f").Build(),
			).Build(),
			NewNode("c").Build(),
			NewNode("d").Build(),
		).Build()

	res := []string{}
	fmt.Println(DepthFirstSearch(res, Tree))
}
