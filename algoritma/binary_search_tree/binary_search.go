package binarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	node  *TreeNode
	level int
}

type BinarySearchTree interface {
	SortedArrayToBST(nums []int) *TreeNode
	PrintBST(treeNode *TreeNode) [][]int
	BSTToSortedArray(bst *TreeNode) []int
}

type meta struct{}

func NewBinarySearchTree() BinarySearchTree {
	return &meta{}
}

func (m *meta) SortedArrayToBST(nums []int) *TreeNode {
	// Time Complexity = O(n)
	// Space Complexity = O(n)
	// n: the total number of elements in the given array
	return constructBst(nums, 0, len(nums)-1)
}

func constructBst(nums []int, startIdx int, endIdx int) *TreeNode {
	if startIdx > endIdx {
		return nil
	}

	midIdx := (startIdx + endIdx) / 2
	bst := &TreeNode{Val: nums[midIdx]}
	bst.Left = constructBst(nums, startIdx, midIdx-1)
	bst.Right = constructBst(nums, midIdx+1, endIdx)

	return bst
}

func (m *meta) PrintBST(root *TreeNode) [][]int {
	nodes := [][]int{}
	queue := []nodeInfo{{node: root, level: 0}}

	for len(queue) > 0 {
		currentNodeInfo := queue[0]
		queue = queue[1:]
		currentNode, currentLevel := currentNodeInfo.node, currentNodeInfo.level
		if len(nodes) == currentLevel {
			nodes = append(nodes, []int{currentNode.Val})
		} else {
			nodes[currentLevel] = append(nodes[currentLevel], currentNode.Val)
		}

		if currentNode.Left != nil {
			queue = append(queue, nodeInfo{node: currentNode.Left, level: currentLevel + 1})
		}
		if currentNode.Right != nil {
			queue = append(queue, nodeInfo{node: currentNode.Right, level: currentLevel + 1})
		}
	}
	return nodes
}

func (m *meta) BSTToSortedArray(bst *TreeNode) []int {
	// Time Complexity = O(n)
	// Space Complexity = O(n)
	// n: the total number of elements in the given array
	order := []int{}
	inOrderTravesal(bst, &order)
	return order
}

func inOrderTravesal(node *TreeNode, order *[]int) {
	if node == nil {
		return
	}

	inOrderTravesal(node.Left, order)
	*order = append(*order, node.Val)
	inOrderTravesal(node.Right, order)
}
