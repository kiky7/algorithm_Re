/*
 * @author:kiky
 * @date: 2021/7/29 10:42 下午
**/

package main

import (
	"algorithm_Re/tree"
)



func main() {

	//var data = "5,3,6,2,4,null,null,1,null"
	var data = "5,1,7,null,null,6,8"
	deser := tree.Constructor()
	root := deser.Deserializebfs(data)
	re := tree.DeleteNode(root,7)


	println(re)
}