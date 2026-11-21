package ds

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	got := ListToSlice(NewList(1, 2, 3))
	if want := []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if got := ListToSlice(NewList()); len(got) != 0 {
		t.Errorf("empty list: got %v", got)
	}
}

func TestTree(t *testing.T) {
	in := []any{3, 9, 20, nil, nil, 15, 7}
	got := TreeToSlice(NewTree(in...))
	if !reflect.DeepEqual(got, in) {
		t.Errorf("got %v, want %v", got, in)
	}
	if NewTree() != nil || NewTree(nil) != nil {
		t.Error("empty tree should be nil")
	}
	if n := FindNode(NewTree(in...), 15); n == nil || n.Val != 15 {
		t.Error("FindNode failed")
	}
}
