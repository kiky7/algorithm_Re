package main

import (
	"strconv"
	"strings"
)

// 代表分隔符的字符
const SEP = ",";
// 代表 null 空指针的字符
const EMPTY = "null";
// 用于拼接字符串
var sb = "";

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

/*序列化  先序*/
func (Codec) serialize(root *TreeNode) string {
	//特殊处理
	if(root == nil){
		return ""
	}
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

/*反序列化 先序*/
func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	//特殊处理
	if(len(sp) == 0){
		return nil
	}

	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}

/* 主函数，将二叉树序列化为字符串 先序*/
func (Codec) serializeD( root *TreeNode) string{
	sb := ""
	serializeH(root, &sb);
	return sb;
}

/* 辅助函数，将二叉树存入 */
func serializeH(root *TreeNode, sb *string) {
	if (root == nil) {
		*sb = *sb + EMPTY + SEP;
		return;
	}

	//前序遍历位置
	*sb = *sb + strconv.Itoa(root.Val) + SEP;

	serializeH(root.Left, sb);
	serializeH(root.Right, sb);
}

/*主函数，将字符串反序列化为二叉树  先序*/
func (Codec) deserializeD(data string) *TreeNode {
	sp := strings.Split(data,SEP)
	if(len(sp) == 0){
		return nil
	}
	re,_ := deserializeH(sp)
	return re
}

/* 辅助函数，节点插入二叉树 */
func deserializeH(data [] string) (*TreeNode,[] string) {
	if data[0] == "null" {
		data = data[1:]
		return nil,data
	}
	val, _ := strconv.Atoi(data[0])
	data = data[1:]

	left,data := deserializeH(data)
	right,data := deserializeH(data)

	return &TreeNode{val, left, right},data
}


//
func main() {
	var data = "1,2,null,4,null,null,3,null,null"
	//ser := Constructor();
	deser := Constructor()
	root := deser.deserializeD(data)
	datas := deser.serializeD(root)
	println(datas)
}