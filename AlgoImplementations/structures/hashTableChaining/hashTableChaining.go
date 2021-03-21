package main

import "fmt"

const StartHashArrSize = 20
const HashPrimeNumber = 13

type ListNode struct {
	nextNode *ListNode
	value    string
}

type List struct {
	root *ListNode
}

func (l *List) Insert(value string) {
	if l.root == nil {
		l.root = &ListNode{
			value: value,
		}
		return
	}

	newNode := &ListNode{
		nextNode: l.root,
		value:    value,
	}

	l.root = newNode
}

func (l *List) Print() {
	for l.root != nil {
		fmt.Print(l.root.value + "->")
		l.root = l.root.nextNode
	}

	fmt.Println("nil")
}

type HashTable struct {
	hashArr []*List
	size    int
}

func InitHashTable() *HashTable {
	hashTable := new(HashTable)

	hashTable.hashArr = make([]*List, StartHashArrSize)

	for i := range hashTable.hashArr {
		hashTable.hashArr[i] = new(List)
	}

	return hashTable
}

func (ht *HashTable) hash(value string) int {
	hash := 0

	for _, ch := range value {
		hash = (hash*HashPrimeNumber + int(ch)) % cap(ht.hashArr)
	}

	return hash
}

func (ht *HashTable) Insert(value string) {
	ht.hashArr[ht.hash(value)].Insert(value)
	ht.size++
}

func (ht *HashTable) Print() {
	for _, list := range ht.hashArr {
		list.Print()
	}
}

func main() {
	hashTable := InitHashTable()

	hashTable.Insert("abc")
	hashTable.Insert("cba")
	hashTable.Insert("a")
	hashTable.Insert("ac")
	hashTable.Insert("bc")
	hashTable.Insert("b")
	hashTable.Insert("c")

	hashTable.Print()
}
