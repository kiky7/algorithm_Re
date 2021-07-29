/*
 * @author:kiky
 * @date: 2021/7/29 10:42 下午
**/

package main

import "algorithm_Re/tree"



func main() {

	preorder := []int{1,2,5,4,6,7,3,8,9}
	inorder := []int{5,2,6,4,7,1,8,3,9}
	root := tree.BuildTree(preorder,inorder)

	println(root)
}