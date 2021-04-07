package main

import (
	"fmt"
)

func anagramCheck(firstStr, secondStr string) int {
	if len(firstStr) != len(secondStr) {
		return 0
	}

	firstStrMap := make(map[rune]int)
	secondStrMap := make(map[rune]int)

	for _, ch := range firstStr {
		firstStrMap[ch]++
	}

	for _, ch := range secondStr {
		secondStrMap[ch]++
	}

	if len(firstStrMap) != len(secondStrMap) {
		return 0
	}

	for keyFirstStr, valueFirstStr := range firstStrMap {
		valueSecondStr, isContains := secondStrMap[keyFirstStr]
		if !isContains || valueSecondStr != valueFirstStr {
			return 0
		}
	}

	return 1
}

func main() {
	var firstStr, secondStr string

	fmt.Scan(&firstStr, &secondStr)

	fmt.Println(anagramCheck(firstStr, secondStr))
}
