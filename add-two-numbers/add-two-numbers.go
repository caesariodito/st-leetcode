/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0}
    currentNode := dummyHead
    carry := 0

    for l1 != nil || l2 != nil {
        var x, y int
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        sum := carry + x + y
        carry = sum / 10
        currentNode.Next = &ListNode{Val: sum % 10}
        currentNode = currentNode.Next
    }

    if carry > 0 {
        currentNode.Next = &ListNode{Val: carry}
    }

    return dummyHead.Next
}