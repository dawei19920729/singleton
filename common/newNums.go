package common

func GetNums(x int) []int {
	var nums []int
	for i := 0; i < x; i++ {
		nums = append(nums, i)
	}
	return nums
}
