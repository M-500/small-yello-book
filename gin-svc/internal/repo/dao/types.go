package dao

import "errors"

// 这里定义一些通用的dao层的错误，比如记录不存在，key重复等

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrKeyDuplicate   = errors.New("key Duplicate")
)
