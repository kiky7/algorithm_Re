/*
 * @author:kiky
 * @date: 2021/8/2 3:15 下午
**/

package tree

import "strconv"

/**
 * @Description: 主函数--寻找重复子树652
 * @param root
 * @return []*TreeNode
 */
func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	comMap := make(map[string]int)
	ret   := make([]*TreeNode,0)

	dnf(root,comMap,&ret)
	return ret
}

func dnf(root *TreeNode,comMap map[string]int, ret *[]*TreeNode) string{
	if root == nil{
		return ""
	}

	left := dnf(root.Left,comMap,ret)
	right := dnf(root.Right,comMap,ret)

	res := strconv.Itoa(root.Val) + "," + left+","+right  //先序

	if v, ok := comMap[res];ok{
		if v == 1{ //重复一次加入结果树
			*ret = append(*ret,root)
		}
		comMap[res]++;
	}else{
		comMap[res] = 1
	}
	return res
}

/*func main() {

	var data = "1,2,null,4,null,null,2,null,4"
	deser := tree.Constructor()
	root := deser.Deserialize(data)
	re := tree.FindDuplicateSubtrees(root)

	println(re)
}*/