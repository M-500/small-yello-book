package utils

import (
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

// 判断字符串是否为空
func IsEmpty(str string) bool {
	return len(str) == 0
}

// 判断字符串是否为空 或者包含空格
func IsBlank(str string) bool {
	return len(str) == 0 || len(str) != len(strings.TrimSpace(str))
}

// Truncate 截断字符串到指定长度,如果超过长度则添加省略号
func Truncate(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-3] + "..."
}

// ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords 计算字符串中的单词数
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// ToSnakeCase 将驼峰命名转换为蛇形命名
func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// Capitalize 将字符串的首字母大写
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// RemoveDuplicateSpaces 移除字符串中的重复空格
func RemoveDuplicateSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// IsNumeric 检查字符串是否只包含数字
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// RegexMatch determines whether key1 matches the pattern of key2 in regular expression.
func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}

// KeyMatch2 determines whether key1 matches the pattern of key2 (similar to RESTful path), key2 can contain a *.
// For example, "/foo/bar" matches "/foo/*", "/resource1" matches "/:resource"
func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

func PathsEqual(path1, path2 string) bool {
	parts1 := strings.Split(path1, "/")
	parts2 := strings.Split(path2, "/")

	if len(parts1) != len(parts2) {
		return false
	}
	for i := 0; i < len(parts1); i++ {
		if strings.HasPrefix(parts2[i], ":") || parts2[i] == "*" {
			continue
		} else if parts1[i] != parts2[i] {
			return false
		}
	}
	return true
}
