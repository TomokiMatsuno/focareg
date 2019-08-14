// trie
package main

import (
	"fmt"
)

// When you use a rune to pick child up from a node.

type Node struct {
	children map[rune]*Node
	is_leaf  bool
}

func (n Node) getChild(r rune) *Node {
	return n.children[r]
}

func (n Node) createChild(r rune) {
	n.children[r] = &Node{
		children: make(map[rune]*Node),
	}
	// n.children[r] = Node{children: make(map[rune]*Node)}
}

func characterToIndex(c rune) int {
	return int(c - '`')
}

// func (n *Node) addWord(word string) {
func (n *Node) addWord(word string) {
	next_node := n
	for _, char := range word {

		v := next_node.getChild(char)

		if v == nil {
			next_node.createChild(char)
		}
		next_node = next_node.children[char]

	}
	next_node.is_leaf = true
}

func (n *Node) searchWord(word string) bool {
	next_node := n
	for _, char := range word {
		if next_node.is_leaf {
			fmt.Print("reached to a leaf")
			return false
		}

		v := next_node.getChild(char)
		// fmt.Println(v)

		if v != nil {
			// next_node.createChild(char_index)
			next_node = next_node.children[char]
		} else {
			fmt.Println("v == nil")
			return false
		}

	}
	return next_node.is_leaf
}
