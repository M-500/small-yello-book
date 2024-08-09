package dao

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrKeyDuplicate   = errors.New("Key Duplicate")
)
