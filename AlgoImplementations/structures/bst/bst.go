package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	left, right, parent *TreeNode
	key                 int
}

type BinaryTree struct {
	root *TreeNode
}

func (bt *BinaryTree) Insert(key int) {
	if bt.root == nil {
		bt.root = &TreeNode{key: key}
		return
	}

	bt.root.Insert(key)
}

func (tn *TreeNode) Insert(key int) {
	if key <= tn.key {
		if tn.left == nil {
			tn.left = &TreeNode{key: key, parent: tn}
		} else {
			tn.left.Insert(key)
		}
	} else {
		if tn.right == nil {
			tn.right = &TreeNode{key: key, parent: tn}
		} else {
			tn.right.Insert(key)
		}
	}
}

func (bt *BinaryTree) InorderWalk() {
	bt.root.InorderWalk()
}

func (tn *TreeNode) InorderWalk() {
	if tn.left != nil {
		tn.left.InorderWalk()
	}

	fmt.Println(tn.key)

	if tn.right != nil {
		tn.right.InorderWalk()
	}
}

func (bt *BinaryTree) Find(key int) *TreeNode {
	it := bt.root

	for it != nil && it.key != key {
		if key <= it.key {
			it = it.left
		} else {
			it = it.right
		}
	}

	return it
}

func (bt *BinaryTree) FetchMin() *TreeNode {
	if bt.root == nil {
		return nil
	}

	return bt.root.FetchMin()
}

func (bt *BinaryTree) FetchMax() *TreeNode {
	if bt.root == nil {
		return nil
	}

	return bt.root.FetchMax()
}

func (tn *TreeNode) FetchMin() *TreeNode {
	for tn.left != nil {
		tn = tn.left
	}

	return tn
}

func (tn *TreeNode) FetchMax() *TreeNode {
	for tn.right != nil {
		tn = tn.right
	}

	return tn
}

func (bt *BinaryTree) Transplant(u, v *TreeNode) {
	if u.parent == nil {
		bt.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	if v != nil {
		v.parent = u.parent
	}
}

func (bt *BinaryTree) Remove(key int) error {
	rmNode := bt.Find(key)

	if rmNode == nil {
		return errors.New("No such node")
	}

	if rmNode.left == nil {
		bt.Transplant(rmNode, rmNode.right)
	} else if rmNode.right == nil {
		bt.Transplant(rmNode, rmNode.left)
	} else {
		nextNode := rmNode.right.FetchMin()

		if nextNode.parent != rmNode {
			bt.Transplant(nextNode, nextNode.right)
			nextNode.right = rmNode.right
			nextNode.right.parent = nextNode
		}

		bt.Transplant(rmNode, nextNode)
		nextNode.left = rmNode.left
		nextNode.left.parent = nextNode
	}

	return nil
}

func (tn *TreeNode) Successor() *TreeNode {
	if tn.right != nil {
		return tn.right.FetchMin()
	}

	it := tn.parent

	for it != nil && tn == it.right {
		tn = it
		it = it.parent
	}

	return it
}

func (tn *TreeNode) Predecessor() *TreeNode {
	if tn.left != nil {
		return tn.left.FetchMax()
	}

	it := tn.parent

	for it != nil && tn == it.left {
		tn = it
		it = it.parent
	}

	return it
}

func main() {
	binaryTree := new(BinaryTree)

	binaryTree.Insert(5)
	binaryTree.Insert(2)
	binaryTree.Insert(7)
	binaryTree.Insert(1)
	binaryTree.Insert(3)
	binaryTree.Insert(6)
	binaryTree.Insert(8)

	binaryTree.InorderWalk()

	findEl := binaryTree.Find(8)

	fmt.Println(findEl)
	fmt.Println(findEl.parent)

	fmt.Println(findEl.Successor())
	fmt.Println(findEl.Predecessor())

	binaryTree.Remove(5)

	binaryTree.InorderWalk()

	fmt.Println(binaryTree.root)

	fmt.Println(binaryTree.FetchMin())
	fmt.Println(binaryTree.FetchMax())
}
