package main

import "fmt"

/*
给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

回文字符串 是正着读和倒过来读一样的字符串。

子字符串 是字符串中的由连续字符组成的一个序列。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
输入：s = "aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
*/

func main() {
	s := "aaa"
	fmt.Printf("countSubstrings = %s", countSubstrings(s))

}
func countSubstrings(s string) int {
	length := len(s)
	count := 0
	for i := 0; i < length; i++ {
		left := i
		right := i
		for left >= 0 && right < length && checkStringIndexEqual(s, left, right) {
			count++
			left--
			right++
		}
	}
	for i := 0; i < length; i++ {
		left := i
		right := i + 1
		for left >= 0 && right < length && checkStringIndexEqual(s, left, right) {
			count++
			left--
			right++
		}
	}
	return count
}
func checkStringIndexEqual(s string, left, right int) bool {
	if s[left] == s[right] {
		return true
	} else {
		return false
	}
}
