package algorithms

import "encoding/json"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (bt *TreeNode) String() string {
	b, _ := json.Marshal(bt)
	return string(b)
}

//func NewIntBTree(vals ...int)  *TreeNode{
//	if len(vals) < 1 {
//		return nil
//	}
//}
