package bfs

import (
	"fmt"
	"testing"
)

func TestRotten(t *testing.T) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Printf("Input: %v\n", grid)
	r := orangesRotting(grid)
	fmt.Printf("r: %v\n", r)
}

func TestLadder(t *testing.T) {
	beginWord, endWord := "hit", "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	// wordList := []string{"hot", "dot", "dog", "lot", "log"}

	// beginWord, endWord := "hot", "dog"
	// wordList := []string{"hot", "dog", "dot"}

	// beginWord, endWord := "a", "c"
	// wordList := []string{"a", "b", "c"}

	// beginWord, endWord := "hot", "dog"
	// wordList := []string{"hot", "dog"}

	fmt.Printf("beginWord: %v, endWord: %v\n", beginWord, endWord)
	fmt.Printf("wordList: %v\n", wordList)
	r := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("r: %v\n", r)
}
