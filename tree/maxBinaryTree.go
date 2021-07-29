/*
 * @author:kiky
 * @date: 2021/7/20 10:42
**/


package tree

/**
 * @Description: 654最大二叉树（根最大，左右）
 * @param nums
 * @return *TreeNode
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if(len(nums) == 0 ){
		return nil
	}
	return build(nums,0,len(nums)-1)
}

/**
 * @Description:
 * @param nums
 * @param begin
 * @param end
 * @return *TreeNode
 */
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


