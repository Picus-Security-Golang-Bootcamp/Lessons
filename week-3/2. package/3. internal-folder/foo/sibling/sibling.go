package sibling

import "github.com/h4yfans/go-patika.dev/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}
