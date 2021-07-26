package main

import (
	"strconv"
	"strings"
)

// 代表分隔符的字符
const SEP = ",";
// 代表 null 空指针的字符
const EMPTY = "-1";
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

func (Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	//特殊处理
	if(sp[0] == "null"){
		return nil
	}

	var build func() *TreeNode
	build = func() *TreeNode {

		if(len(sp) == 0){
			return nil
		}

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

/* 主函数，将二叉树序列化为字符串 */
func (Codec) serializeD( root *TreeNode) string{
	sb := &strings.Builder{}
	serializeH(root, sb);
	return sb;
}

/* 辅助函数，将二叉树存入 */
func serializeH(root *TreeNode, sb *string) {
	if (root == nil) {
		sb = sb + EMPTY + SEP;
		return;
	}

	//前序遍历位置
	sb = sb + strconv.Itoa(root.Val) + SEP;

	serializeH(root.Left, sb);
	serializeH(root.Right, sb);
}

//
func main() {
	var data = "1,2,null,4,null,null,3,null,null"
	//ser := Constructor();
	deser := Constructor()
	root := deser.deserialize(data)
	datas := deser.serializeD(root)
	println(datas)
}