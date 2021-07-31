/*
 * @author:kiky
 * @date: 2021/7/31 4:42 下午
 * @Description: 二叉树操作记录
**/

package tree

/**
 * @Description: 倒转二叉树-226  左右子树交换
 * @param root
 * @return *TreeNode
 */
func invertTree(root *TreeNode) *TreeNode {

	if(root == nil){
		return root;
	}
	//先交换
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	//再递归深入子树
	invertTree(root.Left);
	invertTree(root.Right);

	return root;
}

/**
 * @Description: 主函数 二叉树 填充每个节点的下一个右侧节点指针116
 * @param root
 * @return *Node
 */
func connect(root *Node) *Node {
	if(root == nil){
		return root
	}
	connectTwoNode(root.Left,root.Right)
	return root
}

/**
 * @Description: 辅助函数 连接两个节点-
 * @param node1
 * @param node2
 */
func connectTwoNode(node1 *Node,node2 *Node){
	if(node1 == nil || node2 == nil){
		return
	}

	//连接两个节点
	node1.Next = node2

	//递归连接
	connectTwoNode(node1.Left,node1.Right)
	connectTwoNode(node2.Left,node2.Right)
	connectTwoNode(node1.Right,node2.Left)
}

/**
 * @Description: 二叉树展开为链表-114
 * @param root
 */
func flatten(root *TreeNode)  {
	if(root == nil){
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	// 记录左右子树
	left := root.Left
	right := root.Right

	// 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 将原先的右子树接到当前右子树的末端
	p := root
	for (p.Right != nil) {
		p = p.Right
	}
	p.Right = right
}

