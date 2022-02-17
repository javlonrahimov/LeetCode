package main

func SwapPairs(head *ListNode) *ListNode {

	newHead := &ListNode{
		Next: head,
	}
	prev := newHead
	i := head
	if i == nil {
		return head
	}
	j := head.Next
	for i != nil && j != nil {
		prev.Next = j
		i.Next = j.Next
		j.Next = i
		if i.Next == nil {
			break
		}
		prev, i, j = i, i.Next, i.Next.Next
	}
	return newHead.Next
}
