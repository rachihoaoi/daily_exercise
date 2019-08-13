package main

import "fmt"

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func (t *BinarySearchTree)Insert (v int) {
	t.Root.Insert(v)
}

func (node *TreeNode) Insert (v int) {
	if v < node.Value {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &TreeNode{v, nil,nil}
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &TreeNode{v, nil, nil}
		}
	}

}

func (t * BinarySearchTree) InOrder() []int {
	var res []int
	t.Root.InOrder(&res)
	return res
}

func (t * BinarySearchTree) PreOrder() []int {
	var res []int
	t.Root.PreOrder(&res)
	return res
}

func (t *BinarySearchTree) LastOrder() []int {
	var res []int
	t.Root.LastOrder(&res)
	return res
}

func (node * TreeNode) InOrder(result *[]int) {
	if node.Left != nil {
		node.Left.InOrder(result)
	}
	*result = append(*result, node.Value)
	if node.Right != nil {
		node.Right.InOrder(result)
	}
}

func (node * TreeNode) PreOrder(result *[]int) {
	*result = append(*result, node.Value)
	if node.Left != nil {
		node.Left.PreOrder(result)
	}
	if node.Right != nil {
		node.Right.PreOrder(result)
	}
}

func (node * TreeNode) LastOrder(result *[]int) {

	if node.Left != nil {
		node.Left.LastOrder(result)
	}
	if node.Right != nil {
		node.Right.LastOrder(result)
	}
	*result = append(*result, node.Value)
}

func (t *BinarySearchTree) FindMin() int {
	node := t.Root
	for {
		if node.Left != nil {
			node = node.Left
		}else {
			return node.Value
		}
	}
}

func (t *BinarySearchTree) FindMax() int {
	node := t.Root
	for {
		if node.Right != nil {
			node = node.Right
		}else {
			return node.Value
		}
	}
}

func (t *BinarySearchTree) Contains(v int) bool {
	node := t.Root
	for {
		if node == nil {
			return false
		} else if node.Value == v {
			return true
		} else if node.Value > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}

func (t *BinarySearchTree) Remove(v int) {
	node := t.Root
	for {
		if node == nil {
			return
		} else if node.Value == v {
			if node.Left != nil {
				tmp := node
				for tmp.Left != nil && tmp.Right != nil {
					tmp = tmp.Right
				}
			} else if node.Right != nil {

			}
		} else if node.Value > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}

func main() {
	a  := []int{1,2,3,4,5}
	fmt.Println(a[:copy(a,a[1:])])
	t := BinarySearchTree{&TreeNode{6,nil,nil}}
	t.Insert(4)
	t.Insert(5)
	t.Insert(9)
	t.Insert(7)
	t.Insert(11)
	fmt.Println(t.PreOrder())
	fmt.Println(t.InOrder())
	fmt.Println(t.LastOrder())
	fmt.Println(t.FindMax())
	fmt.Println(t.FindMin())
	fmt.Println(t.Contains(11))
}