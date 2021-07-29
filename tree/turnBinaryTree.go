/*
 * @author:kiky
 * @date: 2021/7/20 10:42
 * @Description: 构建二叉树
**/

package tree

import (
	"strconv"
	"strings"
)

/**
 * @Description: 序列化  先序1
 * @receiver Codec
 * @param root
 * @return string
 */
func (Codec) serialize(root *TreeNode) string {
	//特殊处理
	if(root == nil){
		return ""
	}
	sb := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return sb.String()
}

/**
 * @Description: 反序列化 先序1
 * @receiver Codec
 * @param data
 * @return *TreeNode
 */
func (Codec) deserialize(data string) *TreeNode {

	//特殊处理
	if(len(data) == 0){
		return nil
	}
	sp := strings.Split(data, ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}




/**
 * @Description: 主函数，将二叉树序列化为字符串 先序2
 * @receiver Codec
 * @param root
 * @return string
 */
func (Codec) serializeD( root *TreeNode) string{
	sb := ""
	serializeH(root, &sb);
	return sb;
}

/**
 * @Description: 辅助函数，将二叉树存入
 * @param root
 * @param sb
 */
func serializeH(root *TreeNode, sb *string) {
	if (root == nil) {
		*sb = *sb + EMPTY + SEP;
		return;
	}

	//前序遍历位置
	*sb = *sb + strconv.Itoa(root.Val) + SEP;

	serializeH(root.Left, sb);
	serializeH(root.Right, sb);
}


/**
 * @Description: 主函数，将字符串反序列化为二叉树  先序2
 * @receiver Codec
 * @param data
 * @return *TreeNode
 */
func (Codec) deserializeD(data string) *TreeNode {
	sp := strings.Split(data,SEP)
	if(len(sp) == 0){
		return nil
	}
	re,_ := deserializeH(sp)
	return re
}

/**
 * @Description: 辅助函数，节点插入二叉树
 * @param data
 * @return *TreeNode
 * @return []
 */
func deserializeH(data [] string) (*TreeNode,[] string) {
	if data[0] == "null" {
		data = data[1:]
		return nil,data
	}
	val, _ := strconv.Atoi(data[0])
	data = data[1:]

	left,data := deserializeH(data)
	right,data := deserializeH(data)

	return &TreeNode{val, left, right},data
}


/**
 * @Description: 二叉树序列化为字符串 层序BFS
 * @receiver Codec
 * @param root
 * @return string
 */
func (Codec) serializeBFS(root *TreeNode) string {
	q := []*TreeNode{root} //队列
	res := []string{}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		} else {
			res = append(res, "null")
		}
	}
	return strings.Join(res, ",")
}

/**
 * @Description: 字符串反序列化为二叉树 层序BFS
 * @receiver Codec
 * @param data
 * @return *TreeNode
 */
func (Codec) deserializebfs(data string) *TreeNode {

	if len(data) == 0 {
		return nil
	}
	list := strings.Split(data, ",")

	//获取第一个值，创建树根，入队
	Val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: Val}
	q := []*TreeNode{root}
	cursor := 1

	//cursor下标取数
	for cursor < len(list) {

		//记录队列第一个值、获取数组未取出的两个节点
		node := q[0]
		q = q[1:]
		leftVal := list[cursor]
		rightVal := list[cursor+1]

		//非X则创建子树 并 把子树值入队
		if leftVal != "null" {
			v, _ := strconv.Atoi(leftVal)
			leftNode := &TreeNode{Val: v}
			node.Left = leftNode
			q = append(q, leftNode)
		}
		if rightVal != "null" {
			v, _ := strconv.Atoi(rightVal)
			rightNode := &TreeNode{Val: v}
			node.Right = rightNode
			q = append(q, rightNode)
		}

		//左右因此+2
		cursor += 2
	}
	return root
}

/*
func main() {
	var data = "1,2,null,4,null,null,3,null,null"
	deser := Constructor()
	root := deser.deserialize(data)
	datas := deser.serialize(root)

	var data = "1,2,3,null,null,4,5,null,null,null,null"
	//ser := Constructor();
	deser := Constructor()
	root := deser.deserializebfs(data)
	datas := deser.serializeBFS(root)

	println(datas)
}*/