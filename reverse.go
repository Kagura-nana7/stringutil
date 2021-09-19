//包 stringutil 包含处理字符串的实用函数。
package stringutil

//Reverse 从左到右按运行方式反向返回传入的字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1;i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

