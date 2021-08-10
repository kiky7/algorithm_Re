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


/**
 * @Description: 二叉搜索树中的搜索-700
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。
返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
 * @param root
 * @param val
 * @return *TreeNode
 */
func SearchBST(root *TreeNode, val int) *TreeNode {

	if( root == nil || root.Val == val){
		return root
	}

	if root.Val > val {
		return SearchBST(root.Left,val)
	}

	return SearchBST(root.Right,val)
}



/*
func main() {

	var data = "5,3,6,2,4,null,null,1,null"
	deser := tree.Constructor()
	root := deser.Deserializebfs(data)
	re := tree.ConvertBST(root)

	println(re)
}*/