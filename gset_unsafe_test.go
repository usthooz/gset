package gset

import "testing"

func Test_GsetUnsafe_Add(t *testing.T) {
	sa := New()
	sa.Add("1")
	sa.Add("2")
	t.Logf("len: %d", sa.Len())
}

func Test_GsetUnsafe_Remove(t *testing.T) {
	sa := New()
	sa.Add("1")
	sa.Add("2")
	sa.Remove("1")
	t.Logf("len: %d", sa.Len())
}

func Test_GsetUnsafe_IsEmpty(t *testing.T) {
	sa := New()
	sa.Add("1")
	t.Log("is empty:", sa.IsEmpty())
}

func Test_GsetUnsafe_Merge(t *testing.T) {
	sa := New()
	sa.Add("1")
	sa.Add("2")
	sa1 := New()
	sa1.Add("1")
	sa1.Add("3")
	sa.Merge(sa1)
	t.Logf("sa len: %d", sa.Len())
}

func Test_GsetUnsafe_Clear(t *testing.T) {
	sa := New()
	sa.Add("1")
	sa.Add("2")
	t.Logf("sa len: %d", sa.Len())
	sa.Clear()
	t.Logf("cleared sa len: %d", sa.Len())
}
