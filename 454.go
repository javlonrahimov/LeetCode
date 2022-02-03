package main

func FourSumCount(nums1, nums2, nums3, nums4 []int) int {

	mapp := make(map[int]int)
	l := len(nums1)
	ans := 0

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			mapp[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			ans += mapp[-nums3[i]-nums4[j]]
		}
	}

	return ans
}
