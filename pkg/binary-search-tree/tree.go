package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
	parent *Node
}

func (tree *Node) IsEmpty() bool {
	return tree == nil || (tree.data == 0 && tree.parent == nil && tree.left == nil && tree.right == nil)
}

func (tree *Node) isNull() bool {
	return tree == nil || tree.IsEmpty()
}

func (tree *Node) IsLeaf() bool {
	return tree.parent != nil && tree.left == nil && tree.right == nil
}

func (tree *Node) MinValue() int {
	min := tree.data

	for !tree.left.isNull() {
		min = tree.left.data
		tree = tree.left
	}

	return min
}

func (tree *Node) Insert(data int) {
	if tree.IsEmpty() {
		tree.data = data
	} else {
		if tree.data > data {
			if tree.left.isNull() {
				tree.left = &Node{data, nil, nil, tree}
			} else {
				tree.left.Insert(data)
			}
		} else if tree.data < data {
			if tree.right.isNull() {
				tree.right = &Node{data, nil, nil, tree}
			} else {
				tree.right.Insert(data)
			}
		} else {
			return
		}
	}
}

func (tree *Node) Delete(data int) {
	*tree = deleteNode(data, tree)
}

func deleteNode(data int, tree *Node) Node {
	if tree.IsEmpty() {
		return *tree
	}

	if tree.data > data {
		tree.left.Delete(data)
	} else if tree.data < data {
		tree.right.Delete(data)
	} else {
		if tree.IsLeaf() {
			return Node{}
		} else if tree.left.isNull() {
			return *tree.right
		} else if tree.right.isNull() {
			return *tree.left
		} else {
			tree.data = tree.right.MinValue()
			tree.right.Delete(tree.data)
		}
	}
	return *tree
}

func (tree Node) Search(data int) *Node {
	if tree.IsEmpty() {
		return nil
	}

	if tree.data == data {
		return &tree
	} else if tree.data > data && !tree.left.isNull() {
		return tree.left.Search(data)
	} else if tree.data < data && !tree.right.isNull() {
		return tree.right.Search(data)
	}

	return nil
}

func (tree Node) Print() {
	if tree.IsEmpty() {
		return
	}

	fmt.Println(tree.data)

	if !tree.left.isNull(){
		tree.left.Print()
	}

	if !tree.right.isNull() {
		tree.right.Print()
	}
}


func main() {
	tree := &Node{}


	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(3)


	tree.Delete(5)
	tree.Print()

}