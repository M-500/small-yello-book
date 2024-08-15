package set

// MapSet
// @Description: not thread safe
type MapSet[T comparable] struct {
	m map[T]struct{}
}

func (m *MapSet[T]) Add(key T) {
	m.m[key] = struct{}{}
}
func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

func (s *MapSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Keys return values no guarantee of order
func (s *MapSet[T]) Keys() []T {
	ans := make([]T, 0, len(s.m))
	for key := range s.m {
		ans = append(ans, key)
	}
	return ans
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}
