/*
 * @author:kiky
 * @date: 2021/7/29 10:42 下午
**/

package main

import (
	"algorithm_Re/tree"
)



func main() {

	var data = "1,2,null,4,null,null,2,null,4"
	deser := tree.Constructor()
	root := deser.Deserialize(data)
	re := tree.FindDuplicateSubtrees(root)

	println(re)
}