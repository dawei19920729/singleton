package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(n int) *ListNode {
	result := &ListNode{}
	header := result
	for i := 0; i < n; i++ {
		list := &ListNode{
			Val:  i,
			Next: nil,
		}
		result.Next = list
		result = result.Next
	}
	return header.Next
}
func CreateListByMinMax(min, max int) *ListNode {
	result := &ListNode{}
	header := result
	for i := min; i <= max; i++ {
		list := &ListNode{
			Val:  i,
			Next: nil,
		}
		result.Next = list
		result = result.Next
	}
	return header.Next
}
