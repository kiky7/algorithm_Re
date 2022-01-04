/*
 * @author:kiky
 * @date: 2021/12/27 11:09 上午
**/

package linkedList

import (
	"testing"
)

func TestList(t *testing.T)  {

	list := List{}

	//1.往单链表末尾追加元素2, 3, 4, 5
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	t.Log(list)
	head := list.GetHeadNode()
	list.Append(head)


	re := HasCycleHash(head)

	t.Log(re)
}

func TestListRe(t *testing.T)  {
	list := List{}

	//1.往单链表末尾追加元素2, 3, 4, 5
	list.Append(1)
	list.Append(4)
	list.Append(3)
	list.Append(2)
	list.Append(1)

	t.Log(list)
	head := list.GetHeadNode()
	t.Log(IsPalindromeArr(head))
}

