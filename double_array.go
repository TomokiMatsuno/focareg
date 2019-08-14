// double_array
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
作るもの
- load_double_array
	- load two lines of integers separated by spaces.
	- argment(s)
		- filename
	- returns
		- two arrays of integers
*/

func load_double_array(filename string) ([]int, []int) {
	//
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arrs [2][]int

	array_index := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		// put 0 on the first position of the arrays to mimic 1-indexed array.
		arrs[array_index] = append(arrs[array_index], 0)

		for _, num := range nums {
			i, _ := strconv.Atoi(num)
			arrs[array_index] = append(arrs[array_index], i)
		}
		array_index += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return arrs[0], arrs[1]

}

func load_index_to_word_dictionary(filename string) map[int]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	d := make(map[int]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tmp := strings.Split(line, " ")
		index_as_string, word := tmp[0], tmp[1]
		index, _ := strconv.Atoi(index_as_string)
		d[index] = word

	}

	return d

}

type DoubleArray struct {
	base                     []int
	check                    []int
	index_to_word_dictionary map[int]string
}

func (da *DoubleArray) Initialize(double_array_file, index_to_word_dictionary_file string) {
	da.base, da.check = load_double_array(double_array_file)
	da.index_to_word_dictionary = load_index_to_word_dictionary(index_to_word_dictionary_file)
}

func (da *DoubleArray) CommonPrefixSearch(prefix string) []string {
	/*
		string を受け取って、辞書のなかのそれと一致する prefix を含む語を返す
		終端記号の index を 0 とする
	*/

	var ret []string

	index := 1

	for _, c := range prefix {
		base_value := da.base[index]
		child_node_index := base_value + characterToIndex(c)

		if da.check[child_node_index] != index {
			return ret
		}

		// 終端記号による遷移を持つか確認する
		child_base_value := da.base[child_node_index]

		fmt.Println(characterToIndex(c), base_value, da.check[child_node_index], child_node_index, child_base_value)

		if da.base[child_base_value] < 0 {
			ret = append(ret, da.index_to_word_dictionary[-da.base[child_base_value]])
		}

		index = child_node_index
	}

	return ret
}

// func main() {
// 	fmt.Println("Hello World!")

// }
