package main

import (
	"fmt"
)

func stonesAndJewels(stones, jewels string) int {
	resultMatches := 0

	stonesMap := make(map[rune]struct{})

	for _, stone := range stones {
		stonesMap[stone] = struct{}{}
	}

	for _, jewel := range jewels {
		if _, isContains := stonesMap[jewel]; isContains {
			resultMatches++
		}
	}

	return resultMatches
}

func main() {
	var stones, jewels string

	fmt.Scanf("%s", &stones)
	fmt.Scanf("%s", &jewels)

	fmt.Println(stonesAndJewels(stones, jewels))
}
