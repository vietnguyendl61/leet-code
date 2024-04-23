package main

func main() {}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var buffer int
	head := &ListNode{}
	result := head

	for l1 != nil || l2 != nil || buffer > 0 {
		result.Next = new(ListNode)
		result = result.Next
		if l1 != nil && l2 != nil {
			if sum := l1.Val + l2.Val; sum >= 10 {
				result.Val = sum - 10 + buffer
				buffer = 1
			} else if sum2 := sum + buffer; sum2 >= 10 {
				result.Val = sum2 - 10
				buffer = 1
			} else {
				result.Val = sum + buffer
				buffer = 0
			}

			l1 = l1.Next
			l2 = l2.Next
			continue
		}
		if l1 != nil {
			if sum := l1.Val + buffer; sum >= 10 {
				result.Val = sum - 10
				buffer = 1
			} else {
				result.Val = sum
				buffer = 0
			}
			l1 = l1.Next
			continue
		}
		if l2 != nil {
			if sum := l2.Val + buffer; sum >= 10 {
				result.Val = sum - 10
				buffer = 1
			} else {
				result.Val = sum
				buffer = 0
			}
			l2 = l2.Next
			continue
		}
		result.Val = buffer
		break
	}

	return head.Next
}
