package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Hash  []byte
}

func NewNode(left, right *Node, data []byte) *Node {
	node := Node{Left: left, Right: right}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Hash = hash[:]
	} else {
		prevHashes := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(prevHashes)
		node.Hash = hash[:]
	}
	return &node
}

func NewTree(data [][]byte) *Node {
	var nodes []*Node

	for _, d := range data {
		node := NewNode(nil, nil, d)
		nodes = append(nodes, node)
	}

	for len(nodes) > 1 {
		n := len(nodes)
		m := (n + 1) / 2
		parents := make([]*Node, m)

		for i := 0; i < m; i++ {
			var left, right *Node
			if i*2 < n {
				left = nodes[i*2]
			}
			if i*2+1 < n {
				right = nodes[i*2+1]
			}
			parents[i] = NewNode(left, right, nil)
		}

		nodes = parents
	}

	return nodes[0]
}

func printTree(node *Node, prefix string, isTail bool) {
	fmt.Printf("%s", prefix)
	if isTail {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}

	fmt.Printf("%x\n", node.Hash)

	if node.Left != nil {
		printTree(node.Left, prefix, node.Right == nil)
	}
	if node.Right != nil {
		printTree(node.Right, prefix, true)
	}
}

func main() {
	data := [][]byte{[]byte("hello"), []byte("world"), []byte("merkle"), []byte("tree")}
	root := NewTree(data)
	printTree(root, "", true)
}
