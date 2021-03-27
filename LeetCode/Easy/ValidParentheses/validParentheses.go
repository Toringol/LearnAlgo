/*
* Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
* determine if the input string is valid.
*
* An input string is valid if:
*
* Open brackets must be closed by the same type of brackets.
* Open brackets must be closed in the correct order.
*
* Example 1:
*
* Input: s = "()"
* Output: true
*
* Example 2:
*
* Input: s = "()[]{}"
* Output: true
*
* Example 3:
*
* Input: s = "(]"
* Output: false
*
* Example 4:
*
* Input: s = "([)]"
* Output: false
*
* Example 5:
*
* Input: s = "{[]}"
* Output: true
*
* Constraints:
*
* 1 <= s.length <= 104
* s consists of parentheses only '()[]{}'.
 */

package main

import (
	"errors"
	"fmt"
)

type ListNode struct {
	value rune
	next  *ListNode
}

type Stack struct {
	head *ListNode
}

func (st *Stack) Push(key rune) {
	if st.head == nil {
		st.head = &ListNode{
			value: key,
			next:  nil,
		}

		return
	}

	newHead := &ListNode{
		value: key,
		next:  st.head,
	}

	st.head = newHead
}

func (st *Stack) Pop() (rune, error) {
	if st.head == nil {
		return 0, errors.New("Stack is empty")
	}

	newHead := st.head.next
	value := st.head.value
	st.head = newHead

	return value, nil
}

func isValid(s string) bool {
	stack := new(Stack)
	openBrackets := 0

	for _, bracket := range s {
		switch bracket {
		case '(', '{', '[':
			stack.Push(bracket)
			openBrackets++
		case ')', '}', ']':
			lastOpenBracket, err := stack.Pop()
			if err != nil {
				return false
			}

			if (lastOpenBracket == '(' && bracket != ')') ||
				(lastOpenBracket == '[' && bracket != ']') ||
				(lastOpenBracket == '{' && bracket != '}') {

				return false
			}

			openBrackets--
		}
	}

	if openBrackets != 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("((()(())))"))
}
