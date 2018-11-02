package gset

import (
	"fmt"
	"strings"
	"sync"
)

type GsetSafe struct {
	gset
	l sync.RWMutex
}

// newGsetSafe new gset as thread un safe
func newGsetSafe() *GsetSafe {
	s := &GsetSafe{}
	s.m = make(map[interface{}]bool)
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
