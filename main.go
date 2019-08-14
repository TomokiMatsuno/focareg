//FirstGoProject project main.go
package main

/*
トライ木の実装
1. 辞書型
2. 配列
3. 構造体
4. 関数

*/

import (
	"fmt"
)

// type Node struct {
// 	children, is_leaf string
// }

// When you use a character index to pick child up from a node.

// type Node struct {
// 	children [26]*Node
// 	is_leaf  bool
// }

// func (n Node) getChild(i int) *Node {
// 	return n.children[i]
// }

// func (n Node) createChild(i int) {
// 	n.children[i] = new(Node)
// 	fmt.Println(n.children)
// }

// func (n Node) characterToIndex(c rune) int {
// 	return int(c - 'a')
// }

func main() {
	// n := new(Node)
	n := Node{
		children: make(map[rune]*Node),
	}
	fmt.Println(characterToIndex('a'))
	s := "あいうえお"
	n.addWord(s)
	n.addWord("worry")
	if n.searchWord("あいうえお") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	fmt.Println("Hello World!")
	// base, check := load_double_array("double_array_content.txt")
	// d := load_index_to_word_dictionary("index_to_word_dictionary.txt")
	da := new(DoubleArray)
	da.Initialize("double_array_content.txt", "index_to_word_dictionary.txt")
	fmt.Println(da.CommonPrefixSearch("acb"))

}
