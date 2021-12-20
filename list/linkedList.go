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

/**
 * @Description: 141. 环形链表--快慢指针
 * @param head
 * @return bool
 * 双指针，每次一步，每次两步，会再次相遇证明存在环
 * 退出条件：指针下一个节点为null
 */
func hasCycle(head *ListNode) bool {
	if(head == nil || head.Next == nil){
		return false
	}
	//定义快慢指针
	slow := head
	fast := head.Next
	for slow != fast {
		if(fast == nil || fast.Next == nil){
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}


/**
 * @Description: 141. 环形链表--辅助数据结构--哈希表
 * @param head
 * @return bool
 * 遍历存储，并查看后续遍历的数据是否存在哈希表中
 * 退出条件：指针下一个节点为null
 */
func hasCycleHash(head *ListNode) bool {
	if(head == nil || head.Next == nil){
		return false
	}

	//
	hash := []*
	cur := head
	for cur.Next != nil {

	}


}




/**
 * @Description: 剑指 Offer II 027. 回文链表
 * @param head
 * @return bool
 * 双指针
 */
/*func isPalindrome(head *ListNode) bool {

}*/