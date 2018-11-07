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
	Separate(s Gset)
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

// Union 合并
func Union(set1, set2 Gset, sets ...Gset) Gset {
	u := set1.Copy()
	set2.Each(func(elem interface{}) bool {
		u.Add(elem)
		return true
	})
	for _, set := range sets {
		set.Each(func(elem interface{}) bool {
			u.Add(elem)
			return true
		})
	}
	return u
}

// Difference
func Difference(set1, set2 Gset, sets ...Gset) Gset {
	// tran copy
	s := set1.Copy()
	// remove set2 by set 1
	s.Separate(set2)
	for _, set := range sets {
		s.Separate(set)
	}
	return s
}

// Intersection 交集
func Intersection(set1, set2 Gset, sets ...Gset) Gset {
	all := Union(set1, set2, sets...)
	result := Union(set1, set2, sets...)

	all.Each(func(item interface{}) bool {
		if !set1.Has(item) || !set2.Has(item) {
			result.Remove(item)
		}

		for _, set := range sets {
			if !set.Has(item) {
				result.Remove(item)
			}
		}
		return true
	})
	return result
}
