package gset

type GsetUnSafe struct {
	gset
}

// newGetUnsafe new gset as thread un safe
func newGetUnsafe() *GsetUnSafe {
	s := &GsetUnSafe{}
	s.m = make(map[interface{}]bool)
	var (
		_ Gset = s
	)
	return s
}

// Add add elems
func (s *GsetUnSafe) Add(elems ...interface{}) bool {
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
func (s *GsetUnSafe) Remove(elems ...interface{}) bool {
	if len(elems) == 0 {
		return false
	}
	for _, elem := range elems {
		delete(s.m, elem)
	}
	return true
}

// Len get this set size
func (s *GsetUnSafe) Len() int {
	return len(s.m)
}

// IsEmpty check this set is empty?
func (s *GsetUnSafe) IsEmpty() bool {
	return s.Len() == 0
}

// List set convert to list
func (s *GsetUnSafe) List() []interface{} {
	list := make([]interface{}, 0, len(s.m))
	for elem := range s.m {
		list = append(list, elem)
	}
	return list
}

// Each
func (s *GsetUnSafe) Each(f func(elem interface{}) bool) {
	for elem := range s.m {
		if !f(elem) {
			break
		}
	}
}

// Merge
func (s *GsetUnSafe) Merge(gs Gset) {
	gs.Each(func(elem interface{}) bool {
		s.m[elem] = true
		return true
	})
}
