package set

import "sync"

var _ Set[int] = (*SyncMapSet[int])(nil)

type SyncMapSet[T comparable] struct {
	lock *sync.RWMutex
	m    map[T]struct{}
}

func NewSyncMapSet[T comparable](size int) *SyncMapSet[T] {
	lock := &sync.RWMutex{}
	return &SyncMapSet[T]{
		lock: lock,
		m:    make(map[T]struct{}, size),
	}
}

func (s *SyncMapSet[T]) Add(key T) {
	s.lock.RLock()
	defer s.lock.Unlock()
	s.m[key] = struct{}{}
}

func (s *SyncMapSet[T]) Delete(key T) {
	s.lock.RLock()
	defer s.lock.Unlock()
	delete(s.m, key)
}

func (s *SyncMapSet[T]) Exist(key T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.m[key]
	return ok
}

// Keys return values no guarantee of order
func (s *SyncMapSet[T]) Keys() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	ans := make([]T, 0, len(s.m))
	for key := range s.m {
		ans = append(ans, key)
	}
	return ans
}
