package gset

type (
	SafeType int
)

const (
	ThreadSafe = iota
	ThreadUnSafe
)

type Gset interface {
	Add(elems ...interface{}) bool
	Remove(elems ...interface{}) bool
	Len() int
	IsEmpty() bool
	Has(elems ...interface{}) bool
	String() string
	List() []interface{}
	Each(func(interface{}) bool)
	Merge(s Gset)
	Clear()
	Copy() Gset
}

type gset struct {
	m map[interface{}]bool
}

// New  new gset
func New(safeType SafeType) Gset {
	if safeType == ThreadUnSafe {
		return newGsetUnsafe()
	}
	return newGsetSafe()
}
