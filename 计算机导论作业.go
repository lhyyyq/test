
package main

import (
	"fmt"
	"sort"
)

type IntArray []int

func (self IntArray) Len() int {
	return len(self)
}

func (self IntArray) Less(i, j int) bool {
	return self[i] > self[j]
}

func (self IntArray) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(IntArray(arr))
	fmt.Println(arr)
}