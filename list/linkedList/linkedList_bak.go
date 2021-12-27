/*
 * @author:kiky
 * @date: 2021/12/27 11:38 上午
**/
package linkedList

import "fmt"

type NodeD struct {
	Data      int
	NextPoint *NodeD
	PrePoint  *NodeD
}

type LinkedList struct {
	head    *NodeD
	current *NodeD
	tail    *NodeD
}

func CreatLinkedList()LinkedList {
	data := []int{1, 21, 31, 51, 62, 2, 3, 42, 33, 12, 12}
	link := LinkedList{}
	var currentNode *NodeD
	for i := 0; i < len(data); i++ {
		currentNode = new(NodeD)
		currentNode.Data = data[i]
		insertNode(&link, currentNode)
	}
	showLinkedList(link)
	return link
}

func showLinkedList(link LinkedList) {
	var currentNode *NodeD
	currentNode = link.head
	for {
		fmt.Println("Node:", currentNode.Data)
		if currentNode.NextPoint == nil {
			break
		} else {
			currentNode = currentNode.NextPoint
		}
	}
}

func insertNode(link *LinkedList, node *NodeD) {
	if link.head == nil {
		link.head = node
		link.tail = node
		link.current = node
	} else {
		link.tail.NextPoint = node
		node.PrePoint = link.tail
		link.tail = node
	}
}
