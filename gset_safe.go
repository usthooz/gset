package gset

import (
	"fmt"
	"strings"
	"sync"
)

// GsetSafe this thread safe
type GsetSafe struct {
	gset
	l sync.RWMutex
}

// newGsetSafe new gset as thread un safe
func newGsetSafe(size ...int) *GsetSafe {
	s := &GsetSafe{}
	if len(size) > 0 {
		s.m = make(map[interface{}]bool, size[0])
	} else {
		s.m = make(map[interface{}]bool)
	}
	var (
		_ Gset = s
	)
	return s
}

// Add add elems
func (s *GsetSafe) Add(elems ...interface{}) bool {
	if len(elems) == 0 {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, elem := range elems {
		s.m[elem] = true
	}
	return true
}

// Add add string slice
func (s *GsetSafe) AddStringSlice(elems []string) bool {
	if len(elems) == 0 {
		return false
	}
	for _, elem := range elems {
		if s.m[elem] {
			// exists
			continue
		}
		s.m[elem] = true
	}
	return true
}

// Remove add
func (s *GsetSafe) Remove(elems ...interface{}) bool {
	if len(elems) == 0 {
		return false
	}

	s.l.Lock()
	defer s.l.Unlock()

	for _, elem := range elems {
		delete(s.m, elem)
	}
	return true
}

// Len get this set size
func (s *GsetSafe) Len() int {
	s.l.Lock()
	defer s.l.Unlock()

	l := len(s.m)
	return l
}

// IsEmpty check this set is empty?
func (s *GsetSafe) IsEmpty() bool {
	s.l.Lock()
	defer s.l.Unlock()

	l := s.Len()
	return l == 0
}

// IsEqual
func (s *GsetSafe) IsEqual(t Gset) bool {
	if conv, ok := t.(*GsetSafe); ok {
		conv.l.RLock()
		defer conv.l.RUnlock()
	}
	// return false 长度不相同
	if sameSize := len(s.m) == t.Len(); !sameSize {
		return false
	}
	equal := true
	t.Each(func(item interface{}) bool {
		_, equal = s.m[item]
		return equal
	})
	return equal
}

// IsSubset
func (s *GsetSafe) IsSubset(t Gset) (subset bool) {
	s.l.Lock()
	defer s.l.Unlock()

	subset = true
	t.Each(func(item interface{}) bool {
		_, subset = s.m[item]
		return subset
	})
	return
}

// IsSuperset
func (s *GsetSafe) IsSuperset(t Gset) bool {
	return t.IsSubset(s)
}

// Has
func (s *GsetSafe) Has(elems ...interface{}) bool {
	var (
		has bool
	)
	if len(elems) == 0 {
		// default false
		return has
	}

	s.l.Lock()
	defer s.l.Unlock()

	has = true
	for _, elem := range elems {
		if _, has = s.m[elem]; !has {
			// nothing
			break
		}
	}
	return has
}

// List set convert to list
func (s *GsetSafe) List() []interface{} {
	s.l.Lock()
	defer s.l.Unlock()

	list := make([]interface{}, 0, len(s.m))
	for elem := range s.m {
		list = append(list, elem)
	}
	return list
}

// Each
func (s *GsetSafe) Each(f func(elem interface{}) bool) {
	s.l.Lock()
	defer s.l.Unlock()

	for elem := range s.m {
		if !f(elem) {
			break
		}
	}
}

// Merge
func (s *GsetSafe) Merge(gs Gset) {
	s.l.Lock()
	defer s.l.Unlock()

	gs.Each(func(elem interface{}) bool {
		s.m[elem] = true
		return true
	})
}

// Clear
func (s *GsetSafe) Clear() {
	s.l.Lock()
	defer s.l.Unlock()

	s.m = make(map[interface{}]bool)
}

// String [1,2,3,4]
func (s *GsetSafe) String() string {
	s.l.Lock()
	defer s.l.Unlock()

	t := make([]string, 0, len(s.List()))
	for _, item := range s.List() {
		t = append(t, fmt.Sprintf("%v", item))
	}
	return fmt.Sprintf("[%s]", strings.Join(t, ", "))
}

// Copy returns a new Set
func (s *GsetSafe) Copy() Gset {
	u := newGsetUnsafe()
	for item := range s.m {
		u.Add(item)
	}
	return u
}

// Separate
func (s *GsetSafe) Separate(t Gset) {
	s.Remove(t.List()...)
}
