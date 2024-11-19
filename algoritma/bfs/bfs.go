package bfs

import (
	"slices"
)

/*
Word Ladder (Leetcode #127)

A transformation sequence from word beginWord to word endWord using a dictionary wordList is
a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

- Every adjacent pair of words differs by a single letter.
- Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
- sk == endWord

Given two words, beginWord and endWord, and a dictionary wordList,
return the number of words in the shortest transformation sequence from beginWord to endWord,
or 0 if no such sequence exists.

Example 1:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

Example 2:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// if endWord not in wordList
	if contain := slices.Contains(wordList, endWord); !contain {
		return 0
	}

	// transform wordlist into pattern
	// example :
	// hot will have pattern
	// 1. *ot  2. h*t   3. ho*
	neighbor := make(map[string][]string, 0)
	wordList = append(wordList, beginWord)

	for _, wl := range wordList {
		for i := 0; i < len(wl); i++ {
			pattern := wl[:i] + "#" + wl[i+1:]
			neighbor[pattern] = append(neighbor[pattern], wl)
		}
	}

	// find neighbor after beginWord until endWord
	counter := 1
	queue := []string{beginWord}
	visited := make(map[string]struct{}, 0)

	visited[beginWord] = struct{}{}

	for len(queue) > 0 {
		for range len(queue) {

			// pop queue
			topWord := queue[0]
			queue = queue[1:]

			if topWord == endWord {
				return counter
			}

			for i := 0; i < len(topWord); i++ {
				pattern := topWord[:i] + "#" + topWord[i+1:]

				// find neighbor word possibilities by pattern
				for _, neighWord := range neighbor[pattern] {
					if _, ok := visited[neighWord]; !ok {
						queue = append(queue, neighWord)
						visited[neighWord] = struct{}{}
					}
				}
			}
		}
		counter += 1
	}

	return 0
}

/*
Rotting Oranges ( Leetcode #994 )

You are given an m x n grid where each cell can have one of three values:

0 representing an empty cell,
1 representing a fresh orange, or
2 representing a rotten orange.
Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1.

Example 1:
Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
Output: 4

Example 2:
Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:
Input: grid = [[0,2]]
Output: 0
Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
*/
type MyGrid struct {
	Row    int
	Column int
}

func orangesRotting(grid [][]int) int {
	// Find all rotten orange
	rows, cols := len(grid), len(grid[0])
	fresh, rotten := 0, []MyGrid{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case 1:
				fresh++
			case 2:
				rotten = append(rotten, MyGrid{Row: r, Column: c})
			}
		}
	}

	// specify some corner cases
	if fresh == 0 && len(rotten) == 0 {
		return 0
	} // no apples, just zero iters
	if len(rotten) == 0 {
		return -1
	} // no rotten, it's impossible
	if fresh == 0 {
		return 0
	} // all rotten, just zero iters

	// BFS
	time := 0
	direction := []MyGrid{{Row: 0, Column: 1}, {Row: 0, Column: -1}, {Row: 1, Column: 0}, {Row: -1, Column: 0}}
	for len(rotten) > 0 && fresh > 0 {
		time++
		for range len(rotten) {
			pos := rotten[0]
			rotten = rotten[1:]

			for _, dir := range direction {
				row := dir.Row + pos.Row
				col := dir.Column + pos.Column

				if row >= rows || row < 0 || col >= cols || col < 0 || grid[row][col] != 1 {
					continue
				}

				grid[row][col] = 2
				fresh--
				rotten = append(rotten, MyGrid{Row: row, Column: col})
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return time
}

/*
Binary Tree Level Order Traversal ( Leetcode #102 )

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

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
