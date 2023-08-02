package common

import "math/rand"

// 二叉树：https://wskdsgcf.gitbook.io/mastering-go-zh-cn/05.0/05.3/05.3.1
func main() {

}

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func getRandomTreeNode(deep int) *TreeNode {
	ansTree := TreeNode{
		Value: 0,
	}
	head := ansTree
	for i := 0; i < deep; i++ {

	}
	return &head
}
func CreateRandomTreeNode(random int) *TreeNode {
	if random == 0 {
		return &TreeNode{
			Value: 0,
		}
	}
	return &TreeNode{
		Value: rand.Intn(random),
	}
}
func AddNodeByMaxDeep(value int, node *TreeNode, maxDeep int) {
	for i := 0; i < maxDeep; i++ {
		if addNodeByDefault(value, node, i, maxDeep) == true {
			break
		}
	}

}
func addNodeByDefault(value int, node *TreeNode, deep, maxDeep int) bool {
	if deep >= maxDeep {
		return false
	}
	newNode := &TreeNode{
		Value: value,
	}
	if node.LeftNode == nil {
		node.LeftNode = newNode
		return true
	}
	if node.RightNode == nil {
		node.RightNode = newNode
		return true
	}
	if addNodeByDefault(value, node.LeftNode, deep+1, maxDeep) == true {
		return true
	}
	if addNodeByDefault(value, node.RightNode, deep+1, maxDeep) == true {
		return true
	}
	return false
}
