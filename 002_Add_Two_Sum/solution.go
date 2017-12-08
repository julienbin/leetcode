/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    var head *ListNode
    var rear *ListNode
    for l1 != nil || l2 != nil {
        sum := 0
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        sum += carry
        carry = 0
        if sum >= 10 {
            carry = 1
            sum = sum % 10
        }
        node := &ListNode{
            sum, 
            nil,
        }
        if head == nil {
            head = node
            rear = node
        } else {
            rear.Next = node
            rear = node
        }
    }
    if carry > 0 {
        rear.Next = &ListNode{
            carry,
            nil,
        }
    }
    return head
    
}
