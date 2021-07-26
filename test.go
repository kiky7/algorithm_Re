package main

// 代表分隔符的字符
const SEP = ",";
// 代表 null 空指针的字符
const EMPTY = "#";
// 用于拼接字符串
var sb = "";

type TreeNode struct {
	val string
	left *TreeNode
	right *TreeNode
}


/* 主函数，将二叉树序列化为字符串 */
func serialize( root *TreeNode) string{
 	sb := "";
	serializeH(root, sb);
	return sb;
}

/* 辅助函数，将二叉树存入 StringBuilder */
func serializeH(root *TreeNode, sb string) {
	if (root == nil) {
		sb = sb + EMPTY + SEP;
		return;
	}

	//前序遍历位置
	sb = sb + root.val + SEP;

	serializeH(root.left, sb);
	serializeH(root.right, sb);
}


func main() {
	println(1)
}