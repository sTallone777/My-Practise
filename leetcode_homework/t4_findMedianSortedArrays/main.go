package main

import "fmt"

const UNIT_MIN uint = 0
const UNIT_MAX uint = ^uint(0)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (n1 + n2 + 1) >> 1
	l := 0
	r := n1
	for l < r {
		m1 := l + (r-l)/2
		m2 := k - m1
		if nums1[m1] < nums2[m2-1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}

	m1 := l
	m2 := k - l
	var c1t1, c1t2 uint
	if m1 <= 0 {
		c1t1 = UNIT_MIN
	} else {
		c1t1 = uint(nums1[m1-1])
	}
	if m2 <= 0 {
		c1t2 = UNIT_MIN
	} else {
		c1t2 = uint(nums2[m2-1])
	}
	c1 := max(c1t1, c1t2)
	if (n1+n2)&1 > 0 {
		return float64(c1)
	}

	var c2t1, c2t2 uint
	if m1 >= n1 {
		c2t1 = UNIT_MAX
	} else {
		c2t1 = uint(nums1[m1])
	}
	if m2 >= n2 {
		c2t2 = UNIT_MAX
	} else {
		c2t2 = uint(nums2[m2])
	}
	c2 := min(c2t1, c2t2)
	return float64((c1 + c2) >> 1)
}

func max(n1 uint, n2 uint) uint {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1 uint, n2 uint) uint {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{9, 10, 100}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
