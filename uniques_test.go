package uniques

import (
	"reflect"
	"testing"
)

func TestStringSlice(t *testing.T) {
	var mySlice = []string{"test", "test", "blah", "test"}
	var want = uniqueStrings{"test", "blah"}

	v := StringSlice(mySlice)

	if !reflect.DeepEqual(want, v) {
		t.Fatalf("removeDuplicatesStrSlice(mySlice) = %v, want %v", v, want)
	}
}

func TestStringSliceEmpty(t *testing.T) {
	var mySlice = []string{}
	var want = uniqueStrings{}
	v := StringSlice(mySlice)

	if len(v) > 0 {
		t.Errorf("removeDuplicatesStrSlice(mySlice) = %v, want %v", v, want)
	}
}

func TestStringSliceSort(t *testing.T) {
	var mySlice = []string{"aaa", "AA", "c", "b"}
	var expected = []string{"AA", "aaa", "b", "c"}

	v := StringSlice(mySlice).Sort()
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("StringSlice(mySlice).Sort() = %v, want %v", v, expected)
	}
}

func TestIntSlice(t *testing.T) {
	var mySlice = []int{2, 2, 1, 3, 3, 5, 6, 2}
	var want = uniqueInts{2, 1, 3, 5, 6}

	v := IntSlice(mySlice)

	if !reflect.DeepEqual(want, v) {
		t.Fatalf("IntSlice(mySlice) = %v, want %v", v, want)
	}
}

func TestIntSliceSort(t *testing.T) {
	var mySlice = []int{5, 5, 4, 3, 3, 2, 1}
	var expected = []int{1, 2, 3, 4, 5}

	v := IntSlice(mySlice).Sort()
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("IntSlice(mySlice).Sort() = %v, want %v", v, expected)
	}
}

func TestNewItems(t *testing.T) {
	old := []string{"test1", "test2", "test3"}
	new := []string{"test1", "test2", "test3", "test4", "test5"}
	expected := []string{"test4", "test5"}

	out := NewItems(old, new)
	if !reflect.DeepEqual(out, expected) {
		t.Errorf("NewItems()=%v, want %v", out, expected)
	}
}
