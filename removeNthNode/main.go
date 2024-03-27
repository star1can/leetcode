package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prevNodeToRemoveNode, pos := getNthNodeFromEnd(head, n+1)

	if prevNodeToRemoveNode.Next == nil {
		return nil
	}

	if prevNodeToRemoveNode == head && pos != n {
		return head.Next
	}

	prevNodeToRemoveNode.Next = prevNodeToRemoveNode.Next.Next

	return head
}

func getNthNodeFromEnd(node *ListNode, pos int) (*ListNode, int) {
	if node.Next == nil {
		return node, 1
	}

	next, currPos := getNthNodeFromEnd(node.Next, pos)

	if currPos == pos {
		return next, pos
	}

	return node, currPos + 1
}

func main() {

}
