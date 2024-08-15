package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	// Exist 返回是否存在这个元素
	Exist(key T) bool
	Keys() []T
}
