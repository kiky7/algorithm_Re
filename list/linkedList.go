/*
 * @author:kiky
 * @date: 2021/7/31 4:34 下午
**/

package list

/**
 * @Description: 反转链表206
 * @param head
 * @return *ListNode
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil{
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}