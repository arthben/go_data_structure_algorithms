package dfs

import (
	"fmt"
	"testing"
)

/*
Course Schedule II ( Leetcode #210 )

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where prerequisites[i] = [ai, bi]
indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return the ordering of courses you should take to finish all courses.
If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2:
Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2.
Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:
Input: numCourses = 1, prerequisites = []
Output: [0]
*/
func TestAja(t *testing.T) {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	r := findOrder(numCourses, prerequisites)
	fmt.Printf("r: %v\n", r)
}

type GraphType map[int][]int

func findOrder(numCourses int, prerequisites [][]int) []int {
	ans := make([]int, 0)

	graph := make(GraphType, 0)
	visited := make([]int, numCourses)
	for _, course := range prerequisites {
		graph[course[0]] = append(graph[course[0]], course[1])
	}

	for i := 0; i < numCourses; i++ {
		if ok := findOrderDetail(i, graph, &visited, &ans); !ok {
			return []int{}
		}
	}
	return ans
}

func findOrderDetail(course int, graph GraphType, visited *[]int, ans *[]int) bool {
	isVisit := *visited
	if isVisit[course] == -1 {
		return false
	}
	if isVisit[course] == 1 {
		return true
	}

	isVisit[course] = -1

	for _, neigh := range graph[course] {
		if ok := findOrderDetail(neigh, graph, visited, ans); !ok {
			return false
		}
	}

	isVisit[course] = 1
	*ans = append(*ans, course)
	return true
}

/*
Path SUM II ( LeetCode #113 )

Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.
A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

Example 1:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22

Example 2:
Input: root = [1,2,3], targetSum = 5
Output: []

Example 3:
Input: root = [1,2], targetSum = 0
Output: []
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	currSlice := make([]int, 0)
	pathSumDetail(root, targetSum, currSlice, &ans)
	return ans
}

func pathSumDetail(root *TreeNode, targetSum int, currSlice []int, ans *[][]int) {
	if root == nil {
		return
	}

	currSlice = append(currSlice, root.Val)
	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		t := append([]int{}, currSlice...)
		*ans = append(*ans, t)
		return
	}

	pathSumDetail(root.Left, targetSum, currSlice, ans)
	pathSumDetail(root.Right, targetSum, currSlice, ans)
}

/*
Clone Graph ( Leetcode #133 )

Given a reference of a node in a connected undirected graph.
Return a deep copy (clone) of the graph.

Test case format:
For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.
An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.
The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).


Example 2:
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.


Example 3:
Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.
*/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	dup := make([]*Node, 101)
	traverseCopy(node, dup)
	return dup[node.Val]
}

func traverseCopy(node *Node, dup []*Node) {
	newNode := new(Node)
	newNode.Val = node.Val

	dup[node.Val] = newNode

	for _, neigh := range node.Neighbors {
		if dup[neigh.Val] == nil {
			traverseCopy(neigh, dup)
		}

		newNode.Neighbors = append(newNode.Neighbors, dup[neigh.Val])
	}
}
