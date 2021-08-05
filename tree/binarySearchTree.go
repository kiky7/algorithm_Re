/*
 * @author:kiky
 * @date: 2021/8/4 7:48 下午
**/

package tree

var rank int
var re int


/**
 * @Description:  二叉搜索树中第K小的元素--230
 * @param root
 * @param k
 * @return int
 */
func KthSmallest(root *TreeNode, k int) int {

	if(root == nil){
		return 0
	}

	KthSmallest(root.Left,k)
	rank++
	//比较
	if(rank == k){
		re = root.Val
	}
	KthSmallest(root.Right,k)
	return re
}
