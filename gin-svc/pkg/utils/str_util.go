package utils

import "strings"

// 判断字符串是否为空
func IsEmpty(str string) bool {
	return len(str) == 0
}

// 判断字符串是否为空 或者包含空格
func IsBlank(str string) bool {
	return len(str) == 0 || len(str) != len(strings.TrimSpace(str))
}
