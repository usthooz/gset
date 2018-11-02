package gset

import "sync"

type GsetSafe struct {
	gset
	l sync.RWMutex
}

// newGsetSafe new gset as thread un safe
func newGsetSafe() *GsetUnSafe {
	s := &GsetSafe{}
	s.m = make(map[interface{}]bool)
	var (
		_ Gset = s
	)
	return s
}
