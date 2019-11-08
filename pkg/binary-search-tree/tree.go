package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
	parent *Node
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
				tree.left = &Node{data, nil, nil, tree}
			} else {
				tree.left.Insert(data)
			}
		} else if tree.data < data {
			if tree.right == nil {
				tree.right = &Node{data, nil, nil, tree}
			} else {
				tree.right.Insert(data)
			}
		} else {
			return
		}
	}
}

func (tree Node) Delete(data int) {

}

func (tree Node) Search(data int) *Node {
	if tree.IsEmpty() {
		return nil
	}

	if tree.data == data {
		return &tree
	} else if tree.data > data && tree.left != nil {
		return tree.left.Search(data)
	} else if tree.data < data && tree.right != nil {
		return tree.right.Search(data)
	}

	return nil
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

	if src := tree.Search(2); src != nil {
		fmt.Println("Search Result: ", src.data)
	}

	tree.Print()

}