package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var numA = []int{1, 2, 3, 4, 5}
	var numB = []int{1, 2, 3, 4, 5}
	sum := minimumXORSum(numA, numB)
	fmt.Println(sum)

}
func minimumXORSum(nums1 []int, nums2 []int) int {
	return 0
}
func minimumXORSum1(nums1 []int, nums2 []int) (res int) {

	n := len(nums1)
	cnt := func(cur int) (res int) {

		for k := 0; k < 31; k++ {
			res += 1 & cur
			cur >>= 1
		}
		return
	}

	mask := 1 << n
	dp := make([]int, mask)
	for i := 1; i < mask; i++ {
		dp[i] = 2e9
	}
	for i := 1; i < mask; i++ {
		for j := 0; j < n; j++ {
			if (1<<j)&i > 0 {
				val := dp[(1<<j)^i] + (nums1[cnt(i)-1] ^ nums2[j])
				if val < dp[i] {
					dp[i] = val
				}
			}

		}
	}
	return dp[mask-1]
}
func minimumXORSum2(nums1 []int, nums2 []int) int {
	n, cnt := len(nums1), 1<<len(nums1)
	dp := make([]int, cnt)
	for i := range dp {
		dp[i] = 2e9
	}
	dp[0] = 0
	for c := 0; c < cnt; c++ {
		i := bits.OnesCount(uint(c))
		for j := 0; j < n; j++ {
			if c>>j&1 == 1 {
				dp[c] = min(dp[c], dp[c^1<<j]+(nums1[i-1]^nums2[j]))
			}
		}
	}
	return dp[cnt-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
给你两个整数数组 nums1 和 nums2 ，它们长度都为 n 。

两个数组的 异或值之和 为 (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... + (nums1[n - 1] XOR nums2[n - 1]) （下标从 0 开始）。

比方说，[1,2,3] 和 [3,2,1] 的 异或值之和 等于 (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 + 2 = 4 。
请你将 nums2 中的元素重新排列，使得 异或值之和 最小 。

请你返回重新排列之后的 异或值之和 。
*/
