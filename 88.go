package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, m, n := len(nums1)-1, m-1, n-1

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	for m >= 0 {
		nums1[i] = nums1[m]
		m--
		i--
	}
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
	nums1 = nums1[i+1:]
	return
}
