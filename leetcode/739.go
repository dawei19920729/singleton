package main

import "fmt"

/*
给定一个整数数组  temperatures  ，表示每天的温度，返回一个数组  answer  ，其中  answer[i]  是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用  0 来代替。
输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
*/

func main() {
	arrayTem := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Printf("%s", dailyTemperatures(arrayTem))
}

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}

	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func dailyTemperatures1(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
