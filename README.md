## Gset
[![Build Status](https://travis-ci.org/usthooz/gset.svg?branch=master)](https://travis-ci.org/usthooz/gset)
[![Go Report Card](https://goreportcard.com/badge/github.com/usthooz/gset)](https://goreportcard.com/report/github.com/usthooz/gset)
[![GoDoc](http://godoc.org/github.com/usthooz/gset?status.svg)](http://godoc.org/github.com/usthooz/gset)

Golang实现的集合操作。

### 功能
1. 合并
2. 交集
3. 并集
4. 复制
5. 比较

### Example

```
import "testing"

func Test_GsetUnsafe_Add(t *testing.T) {
	sa := New(ThreadUnSafe)
	sa.Add("1")
	sa.Add("2")
	t.Logf("len: %d", sa.Len())
}

func Test_GsetUnsafe_Remove(t *testing.T) {
	sa := New(ThreadUnSafe)
	sa.Add("1")
	sa.Add("2")
	sa.Remove("1")
	t.Logf("len: %d", sa.Len())
}

func Test_GsetUnsafe_IsEmpty(t *testing.T) {
	sa := New(ThreadUnSafe)
	sa.Add("1")
	t.Log("is empty:", sa.IsEmpty())
}

func Test_GsetUnsafe_Merge(t *testing.T) {
	sa := New(ThreadUnSafe)
	sa.Add("1")
	sa.Add("2")
	sa1 := New(ThreadUnSafe)
	sa1.Add("1")
	sa1.Add("3")
	sa.Merge(sa1)
	t.Logf("sa len: %d", sa.Len())
}

func Test_GsetUnsafe_Clear(t *testing.T) {
	sa := New(ThreadUnSafe)
	sa.Add("1")
	sa.Add("2")
	t.Logf("sa len: %d", sa.Len())
	sa.Clear()
	t.Logf("cleared sa len: %d", sa.Len())
}
```