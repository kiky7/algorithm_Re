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

func (Codec) constructMaximumBinaryTree(nums []int) *TreeNode {
	if(len(nums) == 0 ){
		return nil
	}
	return build(nums,0,len(nums)-1)
}

func build(nums []int,begin int,end int) *TreeNode {

	if(begin > end){
		return nil
	}

	maxVal := nums[begin]
	maxIndex := begin

	for i:=begin;i<=end;i++ {
		if(maxVal < nums[i]){
			maxVal = nums[i]
			maxIndex = i
		}
	}

	root := &TreeNode{Val: maxVal}

	root.Left = build(nums,begin,maxIndex-1)
	root.Right = build(nums,maxIndex+1,end)

	return root
}

