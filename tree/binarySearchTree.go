/*
 * @author:kiky
 * @date: 2021/8/4 7:48 下午
**/

package tree

import "algorithm_Re/lib"

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


/**
 * @Description: n个数字组成的二叉搜索树种树 ---96
 * @param n
 * @return int
 */
func NumTrees(n int) int {
	re := lib.ArrTwoCreate(n+1)
	return count(1,n,re)
}

/**
 * @Description: 辅助函数
 * @param be
 * @param en
 * @param re
 * @return int
 */
func count(be int,en int,re [][]int)  int{

	if(be > en){
		return 1
	}

	if(re[be][en] != 0){
		return re[be][en]
	}

	num := 0
	for i:=be;i<=en;i++ {
		left := count(be,i-1,re)
		right := count(i+1,en,re)
		num += left * right
	}

	re[be][en] = num

	return num
}


/**
 * @Description: n个数字组成的二叉搜索树种树 ---96(简化解法)
 * @param n
 * @return int
 */
func NumTreesTwo(n int) int {
	G := make([]int, n + 1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

/**
 * @Description:  n个数字组成的二叉搜索树种树 ---96(卡特兰数解法)
 * @param n
 * @return int
 */
func NumTreesKT(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2 * i + 1) / (i + 2);
	}
	return C
}




/*func main() {

	//var data = "5,3,6,2,4,null,null,1,null"
	var data = "5,1,7,null,null,6,8"
	deser := tree.Constructor()
	root := deser.Deserializebfs(data)
	re := tree.DeleteNode(root,7)


	println(re)
}*/