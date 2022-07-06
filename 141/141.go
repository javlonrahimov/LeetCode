package main

func hasCycle(head *ListNode) bool {

	s := map[*ListNode]bool{}

	for {
		if head != nil {
			if _, ok := s[head]; ok {
				return true
			}

			s[head] = true

			head = head.Next

		} else {
			break
		}
	}
	return false
}
