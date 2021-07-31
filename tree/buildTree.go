/*
 * @author:kiky
 * @date: 2021/7/29 4:28 下午
**/

package tree

/**
 * @Description: 主函数 前序+中序 构造树
 * @param preorder
 * @param inorder
 * @return *TreeNode
 */
func BuildTreeB(preorder []int, inorder []int) *TreeNode {
	if(len(preorder) == 0 || len(inorder) == 0){
		return nil
	}
	return buildH(preorder,0,len(preorder)-1,inorder,0,len(inorder)-1)
}


/**
 * @Description: 辅助函数
 * @param preorder
 * @param pb
 * @param pe
 * @param inorder
 * @param ib
 * @param ie
 * @return *TreeNode
 */
func buildH(preorder []int,pb int,pe int, inorder []int,ib int,ie int) *TreeNode {
	if(pb > pe || ib > ie){
		return nil
	}

	index := 0
	//获取中序根索引
	for i:=ib;i<=ie;i++ {
		if(inorder[i] == preorder[pb]){
			index = i
		}
	}

	//建立树根
	root := &TreeNode{Val:preorder[pb]}

	leftsize := index - ib
	root.Left = buildH(preorder,pb+1,pb+leftsize,inorder,ib,index-1)
	root.Right = buildH(preorder,pb+leftsize+1,pe,inorder,index+1,ie)

	return root
}


/**
 * @Description: 主函数 中序+后序 构造树
 * @param inorder
 * @param postorder
 * @return *TreeNode
 */
func BuildTreeA(inorder []int, postorder []int) *TreeNode {
	if(len(inorder) ==0 || len(postorder) == 0){
		return nil
	}
	return buildHA(inorder,0,len(inorder)-1,postorder,0,len(postorder)-1)
}

/**
 * @Description: 辅助函数
 * @param inorder
 * @param ib
 * @param ie
 * @param postorder
 * @param pb
 * @param pe
 * @return *TreeNode
 */
func buildHA(inorder []int,ib int,ie int,postorder []int,pb int,pe int) *TreeNode {

	if(ib > ie || pb > pe){
		return nil
	}

	//获取中序根下标
	index := 0
	for i:=ib;i<=ie;i++ {
		if(inorder[i] == postorder[pe]){
			index = i
		}
	}
	//后序最后一个节点为根
	root := &TreeNode{Val: postorder[pe]}
	//右子树长度
	rightLen := ie - index

	root.Left = buildHA(inorder,ib,index-1,postorder,pb,pe-rightLen-1)
	root.Right = buildHA(inorder,index+1,ie,postorder,pe-rightLen,pe-1)
 	return root
}


/*func main() {

	preorder := []int{1,2,5,4,6,7,3,8,9}
	inorder := []int{5,2,6,4,7,1,8,3,9}
	postorder := []int{5,6,7,4,2,8,9,3,1}
	root := BuildTree(preorder,inorder)

	println(root)



	leftsize := index - ib

	root.Left = buildHA(inorder,ib,index-1,postorder,pb,pb+leftsize-1)
	root.Right = buildHA(inorder,index+1,ie,postorder,pb+leftsize,pe-1)

}*/