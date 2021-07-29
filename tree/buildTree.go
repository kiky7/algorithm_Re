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
func BuildTree(preorder []int, inorder []int) *TreeNode {
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



/*func main() {

	preorder := []int{1,2,5,4,6,7,3,8,9}
	inorder := []int{5,2,6,4,7,1,8,3,9}
	root := BuildTree(preorder,inorder)

	println(root)
}*/