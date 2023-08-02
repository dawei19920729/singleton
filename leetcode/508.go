package main

import (
	"fmt"
	"time"
)

/*
给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。

一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。

输入: root = [5,2,-3]
输出: [2,-3,4]
*/
func main() {
	//node := common.CreateRandomTreeNode(0)
	//common.AddNodeByMaxDeep(1, node, 5)
	//common.AddNodeByMaxDeep(2, node, 5)
	//common.AddNodeByMaxDeep(3, node, 5)
	//common.AddNodeByMaxDeep(4, node, 5)
	//common.AddNodeByMaxDeep(5, node, 5)
	//common.AddNodeByMaxDeep(6, node, 5)
	//common.AddNodeByMaxDeep(7, node, 5)
	//fmt.Printf("node : %s", node)
	start := time.Now()
	fmt.Printf("cron start :" + start.Format("2006-01-02-15-04-05"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	return []int{}

}
