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


/**
 * @Description:验证二叉搜索树--98
 * @param root
 * @return bool
 */
func IsValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}


/**
 * @Description: 验证二叉搜索树--辅助函数
 * @param root
 * @param lower
 * @param upper
 * @return bool
 */
func helper(root *TreeNode, lower  *TreeNode, upper  *TreeNode) bool {
	if root == nil {
		return true
	}
	if( lower !=nil && root.Val <= lower.Val || upper !=nil && root.Val >= upper.Val) {
		return false
	}
	return helper(root.Left, lower, root) && helper(root.Right, root, upper)
}


/**
 * @Description:701. 二叉搜索树中的插入操作
 * @param root
 * @param val
 * @return *TreeNode
 */
func InsertIntoBST(root *TreeNode, val int) *TreeNode {

	if (root == nil) {
		return &TreeNode{Val: val}
	}

	if(root.Val > val){
		root.Left = InsertIntoBST(root.Left,val)
	}

	if(root.Val < val){
		root.Right = InsertIntoBST(root.Right,val)
	}

	return root
}


/**
 * @Description:450. 删除二叉搜索树中的节点
 * @param root
 * @param key
 * @return *TreeNode
 */
func DeleteNode(root *TreeNode, key int) *TreeNode {

	if(root == nil){return nil}
	if(root.Val == key){
		//无子节点\有一个子节点
		if (root.Left == nil) {
			return root.Right
		}
		if (root.Right == nil) {
			return root.Left
		}
		//有两个子节点
		min := getMin(root.Right) //获取右子树中最小的节点
		root.Val = min.Val //替换根
		root.Right = DeleteNode(root.Right,min.Val) //删除右子树中最小的节点
	}

	if(root.Val > key){
		root.Left = DeleteNode(root.Left,key)
	}

	if(root.Val < key){
		root.Right = DeleteNode(root.Right,key)
	}
	return root
}

/**
 * @Description: 获取树的最小值
 * @param root
 * @return TreeNode
 */
func getMin(node *TreeNode)  *TreeNode{
	for node.Left != nil {
		node = node.Left
	}
	return node
}


func DdeleteNode(root *TreeNode, key int) *TreeNode {
	// 空树判断
	if root == nil {
		return nil
	}
	// 找到需要被删除的节点
	if root.Val == key {
		// 没有左子树，让right代替root的位置
		if root.Left == nil {
			return root.Right
		}
		// 没有右子树,让left代替root的位置
		if root.Right == nil {
			return root.Left
		}
		// 找后继节点
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Right = DdeleteNode(root.Right, next.Val)
		root.Val = next.Val
		return root
	}
	if root.Val > key {
		root.Left = DdeleteNode(root.Left, key)
	} else {
		root.Right = DdeleteNode(root.Right, key)
	}
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