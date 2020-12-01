package leetcode

import (
	"fmt"
	. "github.com/kozmod/idea-tests/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Maximum Depth of Binary Tree
//Given the root of a binary tree, return its maximum depth.
//A binary tree's maximum depth is the number of nodes along the longest path
//from the root node down to the farthest leaf node.
//        [3]
//		 /	 \
//		[9]  [20]
//           /  \
//         [15] [7]
//Example 1:
//Input: root = [3,9,20,null,null,15,7]
//Output: 3
//
//Example 2:
//Input: root = [1,null,2]
//Output: 2
//
//Example 3:
//Input: root = []
//Output: 0
//
//Example 4:
//Input: root = [0]
//Output: 1

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	testCases := []struct {
		btree *TreeNode
		lvl   int
	}{
		{
			btree: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			lvl: 3,
		},
		{
			btree: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
			},
			lvl: 2,
		},
		{
			btree: &TreeNode{Val: 0},
			lvl:   1,
		},
		{
			btree: nil,
			lvl:   0,
		},
	}
	for i, testCase := range testCases {
		res := maxDepth(testCase.btree)
		assert.Equal(t, testCase.lvl, res,
			fmt.Sprintf("test case # %d, expected # %d, got %d", i, testCase.lvl, res))
	}
}

func maxDepth(root *TreeNode) int {
	res := maxDepthWithLVL(0, root)
	return res
}

func maxDepthWithLVL(l int, root *TreeNode) int {
	if root == nil {
		return l
	}
	l++

	ll := maxDepthWithLVL(l, root.Left)
	rl := maxDepthWithLVL(l, root.Right)
	if ll > rl {
		return ll
	}
	return rl
}
