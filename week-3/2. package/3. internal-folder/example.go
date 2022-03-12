package internal_package_example

import "github.com/h4yfans/go-patika.dev/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}
