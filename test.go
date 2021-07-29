package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if(len(preorder) == 0 || len(inorder) == 0){
		return nil
	}
	return build(preorder,0,len(preorder)-1,inorder,0,len(inorder)-1)
}

func build(preorder []int,pb int,pe int, inorder []int,ib int,ie int) *TreeNode {
	if(pb > pe || ib > ie){
		return nil
	}

	index := 0
	//获取中序根索引
	for i:=ib;i<=ie;i++ {
		if(inorder[i] == preorder[pb]){
			index = i
		}
	}

	//建立树根
	root := &TreeNode{Val:preorder[pb]}

	leftsize := index - ib
	root.Left = build(preorder,pb+1,pb+leftsize,inorder,ib,index-1)
	root.Right = build(preorder,pb+leftsize+1,pe,inorder,index+1,ie)

	return root
}


func main() {

	/*	var data = "1,2,null,4,null,null,3,null,null"
		deser := Constructor()
		root := deser.deserialize(data)
		datas := deser.serialize(root)*/

/*	var data = "1,2,3,null,null,4,5,null,null,null,null"
	//ser := Constructor();
	deser := Constructor()
	root := deser.deserializebfs(data)
	datas := deser.serializeBFS(root)

	println(datas)*/

	preorder := []int{1,2,5,4,6,7,3,8,9}
	inorder := []int{5,2,6,4,7,1,8,3,9}
	root := buildTree(preorder,inorder)

	println(root)
}