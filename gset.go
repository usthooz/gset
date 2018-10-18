package gset

type Gset interface {
	Add(elems ...interface{}) bool
	Remove(elems ...interface{}) bool
	Len() int
	IsEmpty() bool
	List() []interface{}
	Each(func(interface{}) bool)
	Merge(s Gset)
}

type gset struct {
	m map[interface{}]bool
}

// New  new gset
func New() Gset {
	return newGetUnsafe()
}