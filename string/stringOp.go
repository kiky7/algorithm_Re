/*
 * @author:kiky
 * @date: 2021/7/31 4:32 下午
**/

package string

/**
 * @Description: 反转字符串344
 * @param s
 */
func ReverseString(s []byte)  {
	for left,right := 0,len(s)-1;left < right;left++{
		s[left],s[right] = s[right],s[left]
		right--
	}
}
