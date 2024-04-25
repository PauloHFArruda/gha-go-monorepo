package world

import "github.com/elliotchance/pie/v2"

func World() string {
	return "World"
}

func Dummy() bool {
	return pie.Contains([]int{}, 2)
}
