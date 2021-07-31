/*
 * @author:kiky
 * @date: 2021/7/31 4:32 下午
**/

package string

import "strings"


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

/**
 * @Description:反转字符串2-541
 * @param s
 * @param k
 * @return string
 */
func reverseStr(s string, k int) string {
	str_arr := strings.Split(s,"")
	slen := len(str_arr)
	for s:=0;s<slen;s+=2*k{
		i := s
		j := s+k-1
		if (s+k-1 > slen-1){
			j = slen-1
		}
		for i<j {
			tmp := str_arr[i]
			str_arr[i] = str_arr[j]
			i++
			str_arr[j] = tmp
			j--
		}
	}
	return strings.Join(str_arr,"");
}