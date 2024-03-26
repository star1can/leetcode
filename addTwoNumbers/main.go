package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	res := &ListNode{0, nil}
	curr := res

	i := l1
	j := l2
	carry := 0

	for i != nil || j != nil || carry != 0 {
		curr.Next = &ListNode{0, nil}
		curr = curr.Next

		sum := carry

		if i != nil {
			sum += i.Val
			i = i.Next
		}

		if j != nil {
			sum += j.Val
			j = j.Next
		}

		carry = sum / 10
		curr.Val = sum % 10
	}

	return res.Next
}

func main() {

}
