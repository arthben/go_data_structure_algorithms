package binarytreepaths

/*
PreOrderTransversal : root -> Left -> Right
PostOrderTranversal : Left -> Right -> Root
InOrderTranversal   : Left -> Root -> Right
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{Val: v, Left: nil, Right: nil}
}

type meta struct{}

type BinaryTreePath interface {
	PreOrderTransversal(*TreeNode) []int
	PostOrderTransversal(*TreeNode) []int
	InOrderTransversal(*TreeNode) []int
}

func NewBinaryTreePaths() BinaryTreePath {
	return &meta{}
}

func (m *meta) InOrderTransversal(root *TreeNode) []int {
	var ans []int
	inOrderTransversalUtil(root, &ans)
	return ans
}

func inOrderTransversalUtil(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	inOrderTransversalUtil(root.Left, ans)
	*ans = append(*ans, root.Val)
	inOrderTransversalUtil(root.Right, ans)
	return
}

func (m *meta) PreOrderTransversal(root *TreeNode) []int {
	var res []int
	preOrderTransversalUtil(root, &res)
	return res
}

func preOrderTransversalUtil(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	*ans = append(*ans, root.Val)
	preOrderTransversalUtil(root.Left, ans)
	preOrderTransversalUtil(root.Right, ans)
	return
}

func (m *meta) PostOrderTransversal(root *TreeNode) []int {
	var res []int
	postOrderTransversalUtil(root, &res)
	return res
}

func postOrderTransversalUtil(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}

	postOrderTransversalUtil(root.Left, ans)
	postOrderTransversalUtil(root.Right, ans)
	*ans = append(*ans, root.Val)
	return
}
