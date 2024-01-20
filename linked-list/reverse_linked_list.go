package main

import (
	"fmt"
)

func main() {

	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	c := reverseListRecursive(input)
	printList(c)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseListRecursive(head *ListNode) *ListNode {
	return reverse(head, nil)
}

func reverse(cur, prev *ListNode) *ListNode {
	fmt.Printf("next %v prev: %v \v",cur,prev)
	if cur == nil {
		fmt.Printf("lv %v\n",prev)
		return prev
	} 

		next := cur.Next
		cur.Next = prev
		
		return reverse(next, cur)
	
}

func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d ", current.Val)
        current = current.Next
    }
    fmt.Println()
}
