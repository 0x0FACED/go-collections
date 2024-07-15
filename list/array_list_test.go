package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// helper function to compare slices
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
	require.NoError(t, err)
	assert.Equal(t, 1, list.Size())
}

func TestArrayList_Insert(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.Insert(10, 1)
	require.NoError(t, err)
	expected := []int{1, 10, 2, 3}
	assert.True(t, slicesEqual(list.items[:list.Size()], expected))
}

func TestArrayList_RemoveLast(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.RemoveLast()
	require.NoError(t, err)
	expected := []int{1, 2}
	assert.True(t, slicesEqual(list.items[:list.Size()], expected))
}

func TestArrayList_RemoveVal(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	pos, err := list.RemoveVal(2)
	require.NoError(t, err)
	assert.Equal(t, 1, pos)
	expected := []int{1, 3}
	assert.True(t, slicesEqual(list.items[:list.Size()], expected))
}

func TestArrayList_RemoveAt(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.RemoveAt(1)
	require.NoError(t, err)
	expected := []int{1, 3}
	assert.True(t, slicesEqual(list.items[:list.Size()], expected))
}

func TestArrayList_Set(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	err := list.Set(10, 1)
	require.NoError(t, err)
	expected := []int{1, 10, 3}
	assert.True(t, slicesEqual(list.items[:list.Size()], expected))
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	item, err := list.Get(1)
	require.NoError(t, err)
	assert.Equal(t, 2, *item)
}

func TestArrayList_GetPosition(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	pos, err := list.GetPosition(2)
	require.NoError(t, err)
	assert.Equal(t, 1, pos)
}

func TestArrayList_Size(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)

	size := list.Size()
	assert.Equal(t, 2, size)
}

func TestArrayList_Clear(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)

	err := list.Clear()
	require.NoError(t, err)
	assert.Equal(t, 0, list.Size())
}

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	assert.True(t, list.Contains(2))
	assert.False(t, list.Contains(5))
}

type Person struct {
	Name    string
	Surname string
	Age     int
}

func TestArrayList_CustomStruct(t *testing.T) {
	names := []string{"Alex", "Qwerty1", "Tester", "Array", "Queue"}
	surnames := []string{"Alex1", "Qwerty11", "Tester1", "Array1", "Queue1"}

	list := NewArrayList[*Person]()
	for i, name := range names {
		list.Add(&Person{
			Name:    name,
			Surname: surnames[i],
			Age:     i + 10,
		})
	}

	list2 := NewArrayList[*Person]()
	for i, name := range names {
		list2.Add(&Person{
			Name:    name,
			Surname: surnames[i],
			Age:     i,
		})
	}
	assert.NotEqual(t, list, list2)

	list3 := NewArrayList[*Person]()
	for i, name := range names {
		list3.Add(&Person{
			Name:    name,
			Surname: surnames[i],
			Age:     i,
		})
	}
	assert.Equal(t, list3, list2)

	p1, err := list.Get(0)
	assert.NoError(t, err)

	p1_Expected := &Person{
		Name:    "Alex",
		Surname: "Alex1",
		Age:     10,
	}

	assert.Equal(t, &p1_Expected, p1)
}

func TestArrayList_Remove(t *testing.T) {
	list1 := NewArrayList[int]()

	list2 := NewArrayList[int]()

	for i := 0; i < 10000; i++ {
		list1.Add(i)
	}

	assert.Equal(t, 10000, list1.Size())

	for i := 10; i < 10000; i++ {
		list1.RemoveVal(i)
	}

	for i := 0; i < 11; i++ {
		list2.Add(i)
	}

	assert.NotEqual(t, list1, list2)
}

func TestArrayList_RemoveAt2(t *testing.T) {
	list := NewArrayList[int]()

	list.Add(67)
	list.Add(2)
	list.Add(20)
	list.Add(-1234)
	list.Add(542)
	list.Add(9011)

	list.RemoveAt(2)
	list2 := NewArrayList[int]()

	list2.Add(67)
	list2.Add(2)
	list2.Add(-1234)
	list2.Add(542)
	list2.Add(9011)

	assert.Equal(t, list2, list)
}

func TestArrayList_GetPosition2(t *testing.T) {
	list := NewArrayList[int]()
	pseudoRand := []int{13, 6321, 65, 12, 0, 56, 2}
	for i := 0; i < 100; i++ {
		list.Add(i * pseudoRand[i%len(pseudoRand)])
	}
	fmt.Println("list: ", list)

	val, err := list.GetPosition(0)
	assert.NoError(t, err)
	assert.Equal(t, 0, val)
}

func TestArrayList_RemoveEmpty(t *testing.T) {
	list := NewArrayList[int]()
	err := list.RemoveLast()
	assert.Error(t, err)
	assert.Equal(t, "list is empty", err.Error())
}

func comparePersonByName(a, b Person) bool {
	return a.Name < b.Name
}

func TestArrayList_TimSort(t *testing.T) {
	list := NewArrayList[Person]()

	list.Add(Person{Name: "Qerty", Age: 30})
	list.Add(Person{Name: "Bob", Age: 25})
	list.Add(Person{Name: "Alex", Age: 35})

	err := list.Sort(comparePersonByName, TimSort) // 0: TimSort
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for i := 0; i < list.Size(); i++ {
			p, _ := list.Get(i)
			fmt.Println(*p)
		}
	}

	listExpected := NewArrayList[Person]()
	listExpected.Add(Person{Name: "Alex", Age: 35})
	listExpected.Add(Person{Name: "Bob", Age: 25})
	listExpected.Add(Person{Name: "Qerty", Age: 30})

	assert.Equal(t, listExpected, list)
}

func TestArrayList_MergeSort(t *testing.T) {
	list := NewArrayList[Person]()

	list.Add(Person{Name: "Qerty", Age: 30})
	list.Add(Person{Name: "Bob", Age: 25})
	list.Add(Person{Name: "Alex", Age: 35})

	err := list.Sort(comparePersonByName, MergeSort) // 2: MergeSort
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for i := 0; i < list.Size(); i++ {
			p, _ := list.Get(i)
			fmt.Println(*p)
		}
	}

	listExpected := NewArrayList[Person]()
	listExpected.Add(Person{Name: "Alex", Age: 35})
	listExpected.Add(Person{Name: "Bob", Age: 25})
	listExpected.Add(Person{Name: "Qerty", Age: 30})

	assert.Equal(t, listExpected, list)
}

func TestArrayList_QuickSort(t *testing.T) {
	list := NewArrayList[Person]()

	list.Add(Person{Name: "Qerty", Age: 30})
	list.Add(Person{Name: "Bob", Age: 25})
	list.Add(Person{Name: "Alex", Age: 35})

	err := list.Sort(comparePersonByName, QuickSort) // 1: QuickSort
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for i := 0; i < list.Size(); i++ {
			p, _ := list.Get(i)
			fmt.Println(*p)
		}
	}

	listExpected := NewArrayList[Person]()
	listExpected.Add(Person{Name: "Alex", Age: 35})
	listExpected.Add(Person{Name: "Bob", Age: 25})
	listExpected.Add(Person{Name: "Qerty", Age: 30})

	assert.Equal(t, listExpected, list)
}

func TestArrayList_BubbleSort(t *testing.T) {
	list := NewArrayList[Person]()

	list.Add(Person{Name: "Qerty", Age: 30})
	list.Add(Person{Name: "Bob", Age: 25})
	list.Add(Person{Name: "Alex", Age: 35})

	err := list.Sort(comparePersonByName, BubbleSort) // 3: BubbleSort
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for i := 0; i < list.Size(); i++ {
			p, _ := list.Get(i)
			fmt.Println(*p)
		}
	}

	listExpected := NewArrayList[Person]()
	listExpected.Add(Person{Name: "Alex", Age: 35})
	listExpected.Add(Person{Name: "Bob", Age: 25})
	listExpected.Add(Person{Name: "Qerty", Age: 30})

	assert.Equal(t, listExpected, list)
}
