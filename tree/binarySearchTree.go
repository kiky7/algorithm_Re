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

/**
 * @Description: 把二叉搜索树转换为累加树--538
 * @param root
 * @return *TreeNode
 */
func ConvertBST(root *TreeNode) *TreeNode {

	if(root == nil){
		return nil
	}

	ConvertBST(root.Right)

	re += root.Val
	root.Val = re

	ConvertBST(root.Left)

	return root
}

/*
func main() {

	var data = "5,3,6,2,4,null,null,1,null"
	deser := tree.Constructor()
	root := deser.Deserializebfs(data)
	re := tree.ConvertBST(root)

	println(re)
}*/