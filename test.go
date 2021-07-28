package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func constructMaximumBinaryTree(nums []int) *TreeNode {
	if(len(nums) == 0 ){
		return nil
	}
	return build(nums,0,len(nums)-1)
}

func build(nums []int,begin int,end int) *TreeNode {

	if(begin > end){
		return nil
	}

	maxVal := nums[0]
	maxIndex := 0

	for i:=begin;i<end;i++ {
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



func main() {
	data := []int{3,2,1,6,0,5}
	root := constructMaximumBinaryTree(data)

	println(root)
}