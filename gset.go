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
	IsEqual(s Gset) bool
	IsSubset(s Gset) bool
	IsSuperset(s Gset) bool
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
func New(safeType SafeType, size ...int) Gset {
	if safeType == ThreadUnSafe {
		return newGsetUnsafe(size...)
	}
	return newGsetSafe(size...)
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

// StringSlice
func StringSlice(s Gset) []string {
	slice := make([]string, 0)
	for _, item := range s.List() {
		v, ok := item.(string)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// IntSlice
func IntSlice(s Gset) []int {
	slice := make([]int, 0)
	for _, item := range s.List() {
		v, ok := item.(int)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// Int32Slice
func Int32Slice(s Gset) []int32 {
	slice := make([]int32, 0)
	for _, item := range s.List() {
		v, ok := item.(int32)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// Int64Slice
func Int64Slice(s Gset) []int64 {
	slice := make([]int64, 0)
	for _, item := range s.List() {
		v, ok := item.(int64)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// Float32Slice
func Float32Slice(s Gset) []float32 {
	slice := make([]float32, 0)
	for _, item := range s.List() {
		v, ok := item.(float32)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// Float64Slice
func Float64Slice(s Gset) []float64 {
	slice := make([]float64, 0)
	for _, item := range s.List() {
		v, ok := item.(float64)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}

// BoolSlice
func BoolSlice(s Gset) []bool {
	slice := make([]bool, 0)
	for _, item := range s.List() {
		v, ok := item.(bool)
		if !ok {
			continue
		}
		slice = append(slice, v)
	}
	return slice
}
