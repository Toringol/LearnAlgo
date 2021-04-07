package main

import (
	"fmt"
)

func generateBrackets(curBracket string, open, closed, bracketNumber int) {
	if len(curBracket) == 2*bracketNumber {
		fmt.Println(curBracket)
		return
	}

	if open < bracketNumber {
		generateBrackets(curBracket+"(", open+1, closed, bracketNumber)
	}

	if closed < open {
		generateBrackets(curBracket+")", open, closed+1, bracketNumber)
	}
}

func main() {
	var bracketLen int

	fmt.Scan(&bracketLen)

	generateBrackets("", 0, 0, bracketLen)
}
