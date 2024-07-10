package list

import "testing"

// Helper function to compare slices
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList[int]()
	err := list.Add(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
}

func TestArrayList_Insert(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.Insert(10, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 10, 2, 3}
	if !slicesEqual(list.items[:list.Size()], expected) {
		t.Errorf("Expected %v, got %v", expected, list.items[:list.Size()])
	}
}

func TestArrayList_RemoveLast(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.RemoveLast()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 2}
	if !slicesEqual(list.items[:list.Size()], expected) {
		t.Errorf("Expected %v, got %v", expected, list.items[:list.Size()])
	}
}

func TestArrayList_RemoveVal(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	pos, err := list.RemoveVal(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if pos != 1 {
		t.Errorf("Expected position 1, got %d", pos)
	}
	expected := []int{1, 3}
	if !slicesEqual(list.items[:list.Size()], expected) {
		t.Errorf("Expected %v, got %v", expected, list.items[:list.Size()])
	}
}

func TestArrayList_RemoveAt(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.RemoveAt(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 3}
	if !slicesEqual(list.items[:list.Size()], expected) {
		t.Errorf("Expected %v, got %v", expected, list.items[:list.Size()])
	}
}

func TestArrayList_Set(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.Set(10, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := []int{1, 10, 3}
	if !slicesEqual(list.items[:list.Size()], expected) {
		t.Errorf("Expected %v, got %v", expected, list.items[:list.Size()])
	}
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	item, err := list.Get(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if *item != 2 {
		t.Errorf("Expected item 2, got %d", *item)
	}
}

func TestArrayList_GetPosition(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	pos, err := list.GetPosition(2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if pos != 1 {
		t.Errorf("Expected position 1, got %d", pos)
	}
}

func TestArrayList_Size(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)

	size := list.Size()
	if size != 2 {
		t.Errorf("Expected size 2, got %d", size)
	}
}

func TestArrayList_Clear(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)

	err := list.Clear()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0 after clearing, got %d", list.Size())
	}
}
