/*
 * @author:kiky
 * @date: 2021/7/29 10:42 下午
**/

package main

import (
	"algorithm_Re/tree"
)



func main() {

	var data = "5,3,6,2,4,null,null,1,null"
	deser := tree.Constructor()
	root := deser.Deserializebfs(data)
	re := tree.ConvertBST(root)

	println(re)
}