package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {

	return mergeLists(lists, 0, len(lists)-1)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}

}

func mergeLists(lists []*ListNode, start, end int) *ListNode {

	if start > end {
		return nil
	}

	if start == end {
		return lists[start]
	}

	middle := (end + start) / 2

	leftList := mergeLists(lists, start, middle)
	rightList := mergeLists(lists, middle+1, end)

	return mergeTwoLists(leftList, rightList)

}
