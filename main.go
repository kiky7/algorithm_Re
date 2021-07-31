/*
 * @author:kiky
 * @date: 2021/7/29 10:42 下午
**/

package main

import "algorithm_Re/tree"



func main() {

	inorder := []int{5,2,6,4,7,1,8,3,9}
	postorder := []int{5,6,7,4,2,8,9,3,1}
	root := tree.BuildTreeA(inorder,postorder)

	println(root)
}