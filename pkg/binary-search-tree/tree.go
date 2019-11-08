package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func (tree *Node) IsEmpty() bool {
	return tree.data == 0
}

func (tree *Node) Insert(data int) {
	if tree.IsEmpty() {
		tree.data = data
	} else {
		if tree.data > data {
			if tree.left == nil {
				tree.left = &Node{data, nil, nil}
			} else {
				tree.left.Insert(data)
			}
		} else {
			if tree.right == nil {
				tree.right = &Node{data, nil, nil}
			} else {
				tree.right.Insert(data)
			}
		}
	}
}

func (tree Node) Delete(data int) {

}

func (tree Node) Search(data int) {

}

func (tree Node) Print() {
	if tree.IsEmpty() {
		return
	}

	fmt.Println(tree.data)

	if tree.left != nil {
		tree.left.Print()
	}

	if tree.right != nil {
		tree.right.Print()
	}
}


func main() {
	tree := &Node{}

	fmt.Println(tree.IsEmpty())

	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(3)

	tree.Print()

}