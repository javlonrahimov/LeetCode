package main

func detectCycle(head *ListNode) *ListNode {
	s := map[*ListNode]bool{}

	for {
		if head != nil {
			if _, ok := s[head]; ok {
				return head
			}

			s[head] = true

			head = head.Next

		} else {
			break
		}
	}
	return nil
}
