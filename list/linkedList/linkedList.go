/*
 * @author:kiky
 * @date: 2021/12/27 2:23 下午
**/

package linkedList

type Object interface{}

type Node struct {
	Data Object
	Next *Node
}

type List struct {
	headNode *Node //头节点
}

//判断是否为空的单链表
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

//单链表的长度
func (this *List) Length() int {
	cur := this.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

//获取头部节点
func (this *List) GetHeadNode() *Node {
	return this.headNode
}

//从头部添加元素
func (this *List) Add(data Object) {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
}

//从尾部添加元素
func (this *List) Append(data Object) {
	node := &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

//在指定位置添加元素
func (this *List) Insert(index int, data Object) {
	if index < 0 {
		this.Add(data)
	} else if index > this.Length() {
		this.Append(data)
	} else {
		pre := this.headNode
		count := 0
		for count < (index - 1) {
			pre = pre.Next
			count++
		}
		//当循环退出后，pre指向index -1的位置
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除指定元素
func (this *List) Remove(data Object) {
	pre := this.headNode

	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

//删除指定位置的元素
func (this *List) RemoveAtIndex(index int) {
	pre := this.headNode
	if index <= 0 {
		this.headNode = pre.Next
	} else if index >= this.Length() {
		//报错 err
	} else {
		count := 0 //index = 3
		for count != (index-1) && pre.Next != nil {
			count++        //2
			pre = pre.Next //2
		}
		pre.Next = pre.Next.Next
	}
}

//是否包含某个元素
func (this *List) Contain(data Object) bool {
	cur := this.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}