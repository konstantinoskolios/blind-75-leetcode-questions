package main

import "fmt"

func main() {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	input2 := &ListNode{Val: 1, Next: nil}
	h1 := removeNthFromEnd(input,2)
	h2 := removeNthFromEnd(input2,1)
	printList(h1)
	printList(h2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    dummy := &ListNode{Val: 0, Next: head}
    fast := dummy
    slow := dummy

    for i := 0; i <= n; i++ {
		fast = fast.Next
    }
	
    for fast != nil {
		fast = fast.Next
        slow = slow.Next
    }
	
    slow.Next = slow.Next.Next

    return dummy.Next
}

func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d ", current.Val)
        current = current.Next
    }
    fmt.Println()
}
