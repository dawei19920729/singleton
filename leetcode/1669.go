package main

import (
	"fmt"
	"singleton/common"
)

func main() {
	list1 := common.CreateList(7)
	List2 := common.CreateListByMinMax(6, 10)
	mergeInBetween(list1, 1, 4, List2)
	fmt.Println("www")

}

func mergeInBetween(list1 *common.ListNode, a int, b int, list2 *common.ListNode) *common.ListNode {
	header := &common.ListNode{
		Val:  0,
		Next: list1,
	}
	ListIndexA := header
	ListIndexB := list1
	for i := 0; i < b; i++ {
		if i < a {
			ListIndexA = ListIndexA.Next
		}
		ListIndexB = ListIndexB.Next
	}
	ListIndexA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = ListIndexB.Next
	return header.Next
}
