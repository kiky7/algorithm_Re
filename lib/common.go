/*
 * @author:kiky
 * @date: 2021/8/25 7:42 下午
**/

package lib

/**
 * @Description: 获取n*n二维数组，并且初始化值为0
 * @param n
 * @return [][]int
 */
func ArrTwoCreate(n int) [][]int {
	re := make([][]int,0)
	for i:=0;i<n;i++ {
		ree := make([]int,0)
		for j:=0;j<n;j++ {
			ree = append(ree,0)
		}
		re = append(re,ree)
	}
	return re
}