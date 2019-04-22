package main

import (
	"fmt"
)

func  main() {
	bestfinish := bestLeagureFinishes(13,12,435,45231,432,4)
	fmt.Println(bestfinish)
}

func bestLeagureFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}