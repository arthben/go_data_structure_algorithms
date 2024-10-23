package reverselinkedlist

import (
	"fmt"
	"slices"
	"testing"

	linkedlist "github.com/arthben/go_data_structure_algorithms/data_structure/linked_list"
)

func TestWae(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"

	n := len(s1)
	c1, c2 := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		c1[s1[i]-'a']++
		c2[s2[i]-'a']++
	}
	fmt.Printf("s1: %v, pattern: %v\n", s1, c1)
	fmt.Printf("s2: %v, map: %v\n", s2, c2)

	if c1 == c2 {
		fmt.Println("True")
	}

	fmt.Println("=================================")

	for i := n; i < len(s2); i++ {
		idx := s2[i] - 'a'
		firstIdx := s2[i-n] - 'a'

		c2[idx]++
		c2[firstIdx]--
		fmt.Printf("current: %v, prev: %v, map: %v\n", string(s2[i]), string(s2[i-n]), c2)

		if c1 == c2 {
			fmt.Println("True")
		}
	}

	fmt.Println("Salah")
}

func TestSwap(t *testing.T) {
	ll := setupLinkedList()
	fmt.Println("Before Reverse ")
	ll.Print()

	fmt.Println("After Reverse")
	final := swapPairs(ll.LinkededListHead())
	PrintMe(final)
}

func swapPairs(head *linkedlist.Node) *linkedlist.Node {
	if head == nil || head.Next == nil {
		return head
	}
	first, second, third := head, head.Next, swapPairs(head.Next.Next)
	first.Next = third
	second.Next = first
	// PrintMe(second)
	return second
}

func TestLink(t *testing.T) {
	ll := setupLinkedList()
	fmt.Println("Before Reverse ")
	ll.Print()

	fmt.Println("After Reverse")
	left := 2
	right := 4
	head := ll.LinkededListHead()
	// reverseList(head)
	reverseBetween(head, left, right)
}

func reverseBetween(head *linkedlist.Node, left int, right int) {

	// this is dummy node which connected next node to head
	// the dummy node will be in fron of head node
	final := &linkedlist.Node{
		Val:  0,
		Next: head,
	}
	// attach previous node to dummy node
	prev := final

	// track most left node which become last previous node
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// current will hold node at left position
	current := prev.Next

	// flip the node from left to right
	for i := 0; i < right-left; i++ {
		tmp := current.Next
		current.Next = tmp.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}

	PrintMe(final.Next)
}

func setupLinkedList() linkedlist.LinkedList {
	link := linkedlist.NewLinkedList()
	for i := 1; i <= 4; i++ {
		link.AddLast(i)
	}
	return link
}

func PrintMe(head *linkedlist.Node) {
	current := head
	for current != nil {
		fmt.Printf("current.val: %v\n", current.Val)
		current = current.Next
	}
}

func reverseList(head *linkedlist.Node) {
	var (
		previous *linkedlist.Node
		temp     *linkedlist.Node
	)
	current := head

	for current != nil {
		temp = current.Next
		current.Next = previous
		previous = current
		current = temp
	}

	PrintMe(previous)
}

func reverseMe(head *linkedlist.Node) *linkedlist.Node {
	if head == nil || head.Next == nil {
		return head
	}

	head, head.Next = head.Next, reverseMe(head.Next.Next)
	return head
}

func TestReverse(t *testing.T) {
	ll := setupLinkedList()
	fmt.Println("Before Reverse ")
	ll.Print()

	fmt.Println("After Reverse")
	head := ll.LinkededListHead()
	// reverseList(head)
	final := reverseMe(head)
	PrintMe(final)
}

/**
 * Definition for singly-linked list.
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}

// 	var (
// 		previous *ListNode
// 		temp     *ListNode
// 	)
// 	current := head

// 	for current != nil {
// 		temp = current.Next
// 		current.Next = previous
// 		previous = current
// 		current = temp
// 	}

// 	return previous
// }

/*
Divide players into Teams of equals skill ( Leetcode #2491 )

You are given a positive integer array skill of even length n where skill[i] denotes the skill of the ith player.
Divide the players into n / 2 teams of size 2 such that the total skill of each team is equal.

The chemistry of a team is equal to the product of the skills of the players on that team.

Return the sum of the chemistry of all the teams, or return -1 if there is no way to divide the players into teams such that the total skill of each team is equal.

Example 1:

Input: skill = [3,2,5,1,3,4]
Output: 22
Explanation:
Divide the players into the following teams: (1, 5), (2, 4), (3, 3), where each team has a total skill of 6.
The sum of the chemistry of all the teams is: 1 * 5 + 2 * 4 + 3 * 3 = 5 + 8 + 9 = 22.
Example 2:

Input: skill = [3,4]
Output: 12
Explanation:
The two players form a team with a total skill of 7.
The chemistry of the team is 3 * 4 = 12.
Example 3:

Input: skill = [1,1,2,3]
Output: -1
Explanation:
There is no way to divide the players into teams such that the total skill of each team is equal.
*/
func TestPlayers(t *testing.T) {
	skill := []int{3, 2, 5, 1, 3, 4}
	// skill := []int{5, 1, 3, 6}
	res := dividePlayers(skill)
	fmt.Printf("res: %v\n", res)
}

func dividePlayers(skill []int) int64 {
	slices.Sort(skill)
	n := len(skill)
	sum := 0
	totalSkill := skill[0] + skill[n-1]
	for i := 0; i < n/2; i++ {
		if skill[i]+skill[n-i-1] != totalSkill {
			return -1
		}
		sum += skill[i] * skill[n-i-1]
	}

	return int64(sum)
}
