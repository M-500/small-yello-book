package middleware

import (
	"fmt"
	"testing"
	"time"
)

func TestCorsMdl(t *testing.T) {
	timestamp := time.Now().Unix()

	// 转换为字符串并提取前 6 位
	low6 := timestamp % 100000

	fmt.Println("当前 Unix 时间戳：", timestamp)
	fmt.Println("区分度最高的 6 位：", low6)
	//t.Log(unix/60/60/24/365 + 1970)
}
