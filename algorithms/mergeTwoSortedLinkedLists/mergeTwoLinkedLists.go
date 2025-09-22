package main

import (
	"fmt"
	"slices"
)

func main() {
	list1 := &node{
		data: 1,
		next: &node{
			data: 3,
			next: &node{
				data: 4,
				next: nil,
			},
		},
	}

	list2 := &node{
		data: 1,
		next: &node{
			data: 7,
			next: &node{
				data: 9,
				next: nil,
			},
		},
	}

	ls := sortedMerge(list1, list2)

	printList(ls)
}

type node struct {
	data int
	next *node
}

func sortedMerge(list1 *node, list2 *node) *node {
	arr := make([]int, 0)

	for list1 != nil {
		arr = append(arr, list1.data)
		list1 = list1.next
	}

	for list2 != nil {
		arr = append(arr, list2.data)
		list2 = list2.next
	}

	slices.Sort(arr)

	dummy := &node{
		data: -1,
		next: nil,
	}
	curr := dummy

	for _, v := range arr {
		curr.next = &node{
			data: v,
			next: nil,
		}

		curr = curr.next
	}

	return dummy.next
}

func printList(list1 *node) {
	for list1 != nil {
		fmt.Println(list1.data)
		list1 = list1.next
	}
}
