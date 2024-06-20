package tree

import (
	"fmt"
	"math"

	"github.com/foureyez/goleetcode/utils"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150

func maxDepth(root *TreeNode) int {
	return int(maxDepthIterative(root))
	// return int(maxDepthRecursive(0, root))
}

func maxDepthRecursive(level float64, node *TreeNode) float64 {
	if node == nil {
		return level
	}

	return math.Max(maxDepthRecursive(level+1, node.Left), maxDepthRecursive(level+1, node.Right))
}

func maxDepthIterative(node *TreeNode) int {
	if node == nil {
		return 0
	}

	currLevelNodes := utils.Queue[TreeNode]{}
	currLevelNodes.Add(node)
	maxDepth := 0
	for currLevelNodes.Size() > 0 {
		size := currLevelNodes.Size()
		for i := 0; i < size; i++ {
			currNode := currLevelNodes.Remove()
			if currNode.Left != nil {
				currLevelNodes.Add(currNode.Left)
			}

			if currNode.Right != nil {
				currLevelNodes.Add(currNode.Right)
			}
		}
		maxDepth++
	}
	return maxDepth
}
