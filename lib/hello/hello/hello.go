package hello

import "github.com/samber/lo"

func Hello() string {
	return "Hello"
}

func Dummy() bool {
	return lo.Contains([]int{}, 2)
}
