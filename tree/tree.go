/*
 * @author:kiky
 * @date: 2021/7/20 10:42
**/

package tree

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

