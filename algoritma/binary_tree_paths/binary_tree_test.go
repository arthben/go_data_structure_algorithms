package binarytreepaths

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestInOrder(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)
	root.Left.Left = NewTreeNode(4)
	root.Left.Right = NewTreeNode(5)

	btp := NewBinaryTreePaths()
	inOrder := btp.InOrderTransversal(root)
	fmt.Printf("inOrder: %v\n", inOrder)

	preOrder := btp.PreOrderTransversal(root)
	fmt.Printf("preOrder: %v\n", preOrder)

	postOrder := btp.PostOrderTransversal(root)
	fmt.Printf("postOrder: %v\n", postOrder)

}

/*
Binary Tree Maximum Path ( Leetcode #124 )

A path in a binary tree is a sequence of nodes where each pair of
adjacent nodes in the sequence has an edge connecting them.
A node can only appear in the sequence at most once.
Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.
Given the root of a binary tree, return the maximum path sum of any non-empty path.

Example 1:
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

Example 2:
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
*/
func TestMaxPathSum(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Right = NewTreeNode(3)

	ans := maxPathSum(root)
	fmt.Printf("ans: %v\n", ans)
}

func maxPathSum(root *TreeNode) int {
	ans := -1 << 63
	maxPathSumDetail(root, &ans)
	return ans
}

func maxPathSumDetail(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	leftSum := maxPathSumDetail(root.Left, ans)
	rightSum := maxPathSumDetail(root.Right, ans)
	temp := leftSum + root.Val + rightSum
	*ans = max(*ans, temp)

	return max(max(leftSum+root.Val, rightSum+root.Val), 0)
}

/*
Binary Tree Paths ( Leetcode #257 )

Given the root of a binary tree, return all root-to-leaf paths in any order.
A leaf is a node with no children.

Example 1:
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:
Input: root = [1]
Output: ["1"]
*/
func TestBTP(t *testing.T) {
	root := NewTreeNode(1)
	root.Left = NewTreeNode(2)
	root.Left.Right = NewTreeNode(5)
	root.Right = NewTreeNode(3)

	ans := binaryTreePaths(root)
	fmt.Printf("ans: %v\n", ans)
}

func binaryTreePaths(root *TreeNode) []string {
	temp := make([]string, 0)
	ans := make([]string, 0)

	binaryTreePathsDetail(root, temp, &ans)
	return ans
}

func binaryTreePathsDetail(root *TreeNode, temp []string, ans *[]string) {
	if root == nil {
		return
	}

	temp = append(temp, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*ans = append(*ans, strings.Join(temp, "->"))
	} else {
		binaryTreePathsDetail(root.Left, temp, ans)
		binaryTreePathsDetail(root.Right, temp, ans)
	}

	// clear temp
	temp = temp[:len(temp)-1]
	return
}

/*
Kth Smallest element in a BST ( Leetcode #230 )

Given the root of a binary search tree, and an integer k,
return the kth smallest value (1-indexed) of all the values
of the nodes in the tree.

Example 1:
Input: root = [3,1,4,null,2], k = 1
Output: 1

Example 2:
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
*/
func TestKth(t *testing.T) {
	// [5,3,6,2,4,null,null,1]
	root := NewTreeNode(5)
	root.Left = NewTreeNode(3)
	root.Left.Left = NewTreeNode(2)
	root.Left.Right = NewTreeNode(4)
	root.Left.Left.Left = NewTreeNode(1)

	root.Right = NewTreeNode(6)

	k := 3
	ans := kthSmallest(root, k)
	fmt.Printf("ans: %v\n", ans)
}

func kthSmallest(root *TreeNode, k int) int {
	var ans []int
	kthSmallestDetail(root, &ans)
	return ans[k-1]
}

func kthSmallestDetail(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	kthSmallestDetail(root.Left, ans)
	*ans = append(*ans, root.Val)
	kthSmallestDetail(root.Right, ans)
	return
}
