/*
 * @author:kiky
 * @date: 2021/7/31 4:34 下午
 * 链表常用工具 curr 当前节点  pre前节点  next下一个节点
**/

package linkedList

import (
	"algorithm_Re/list/list"
)

/**
 * @Description: 反转链表206
 * @param head
 * @return *ListNode
 */
func reverseList(head *Node) *Node {
	var pre *Node
	curr := head
	for curr != nil{
		next := curr.Next //避免丢失应该先记录
		curr.Next = pre //链表转向
		pre = curr //移动前节点
		curr = next //移动当前节点
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
func hasCycle(head *list.ListNode) bool {
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
 * 循环：遍历存储，并查看后续遍历的数据是否存在哈希表中
 * 退出条件：指针下一个节点为null
 * 如果链表中存在环，则返回 true 。 否则，返回 false。
 */
func HasCycleHash(head *Node) bool {
	if(head == nil || head.Next == nil){
		return false
	}

	hash := map[*Node]struct{}{}
	cur := head

	for cur != nil {
		if _,ok:= hash[cur]; ok{  //不存在则保存
			return true
		}
		hash[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}




/**
 * @Description: 剑指 Offer II 027. 回文链表
 * @param head
 * @return bool
 * 双指针
 */
func IsPalindrome(head *Node) bool {
	reList := reverseList(head)
	if head == reList {
		return true
	}
	return false
}


func IsPalindromeArr(head *Node) bool {
	arr := []Object{}
	cur := head
	for cur != nil{
		arr = append(arr,cur.Data)
		cur = cur.Next
	}
	num:= len(arr)
	for key,_  := range arr[:num/2]{
		if arr[key] != arr[num-1-key] {
			return false
		}
	}
	return true
}