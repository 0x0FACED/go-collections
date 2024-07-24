# Singly-Linked List

## Abstract

This document contains all the necessary information about the Singly-Linked List **(next SLL)**. 

This data structure, methods of insertion, deletion, search, and others are described here. As well as a description of the specific implementation provided in this library.

## Table of Contents

- [Description](#description)
    - [Node Structure](#node-structure)
- [Basic Operations](#basic-operations)
    - [Insertion](#insertion)
    - [Deleteion](#deletion)
    - [Searching](#searching)
    - [Traversal](#traversal)
    - [Size](#size)
- [Implementation of SLL in Golang](#implementation)
- [Usage](#usage)

## Description

**Singly-Linked List** - the simplest type of `Linked List`, where each node contains some data and a reference to the next `node`. They can only be traversed in a single direction - from the head (the first node) to the tail (the last node) - that's why this data structure called **SLL** (Singly-Linked List).

![Singly-Linked List](./images/singly_linked_list_g.svg)

The picture above shows a SLL:
1. `HEAD` - pointer to the first item in the list
2. `TAIL` is a pointer to the last element (it is not always added to the implementation, but it simplifies some operations, this will be shown below)
3. `Val` and `Next` represent a `Node`
4. `Val` is the value stored in this node 
5. `Next` is a pointer to the next node

`TAIL` should point to a new node that does not exist yet, that is, to `NULL`,

### Node Structure

For example, we will take the value in the node as an int.
The node structure can be represented as follows:

```go
type Node struct {
    Val int
    Next *Node
}
```

Each block in the first picture represents such a structure. It is worth noting that the list is unidirectional, that is, we cannot move in the opposite direction - from TAIL to HEAD.

## Basic Operations

With such a list, you can come up with many operations, but it is worth to describe the following operations:

1. **Insertion**
2. **Deletion**
3. **Searching**
4. **Traversal**
5. **Size**

### Insertion

The insertion into the list can be in different variations:

- Insert at the start
- Insert in the middle
- Insert at the end

Here is the sample code to implement the insertion written in Golang:

```go
type Node struct {
    Val int
    Next *Node
}

type List struct {
    Head *Node
    Tail *Node
}

// Visualization - Insert at the start
//                       HEAD                    TAIL
//                        |                       |
// 1.   we have list: [1 | next]->[2 | next]->[3 | next]->NULL
// 2.   call insert at start with item = -1
// 3.   create new node:
//      newNode = [-1 | NULL] (newNode.Next = null right now)
//                 list = [1 | next]->[2 | next]->[3 | next]->NULL
//             no connection between list and newNode
//      newNode = [-1 | NULL] 
// 4.   link newNode's Next to head: 
//                      HEAD                   TAIL
//                       |                      |
//      [-1 | next]->[1 | next]->[2 | next]->[3 | next]->NULL
//      But our head still at old position
// 5.   move head link to new node:
//          HEAD                                TAIL
//           |                                   |
//      [-1 | next]->[1 | next]->[2 | next]->[3 | next]->NULL 
func (l *List)InsertAtStart(item int) {
    // Create new node
    newNode := &Node{Val: item}
    // link next of new node to head
    newNode.Next = l.Head
    // move head to new node
    l.Head = newNode
}

// Insert in the middle
func (l *List)InsertAtPos(item int, pos int) {
    // create new node
    newNode := &Node{Val: item}
    // counter to pos
    counter := 0
    // dummy for traversal
    dummy := l.Head
    // while next of dummy != null OR counter < pos-1
    for dummy.Next != nil || counter < pos-1 {
        dummy = dummy.Next
        counter++
    }
    // if true -> pos > size of list
    if counter < pos-1 {
        return
    }
    // we stopped in front of the element that the new one will point to
    newNode.Next = dummy.Next
    dummy.Next = newNode
}

// Thats why we use Tail -> fast insert at the end
// if there was no tail, we would have to traverse from head to last element and do insert
func (l *List)InsertAtEnd(item int) {
    newNode := &Node{Val: item}
    l.Tail.Next = newNode
    l.Tail = newNode
}
```

### Deletion

Same as [Insertion](#insertion) deletion have 3 main variations:

- Delete from the start
- Delete from the middle
- Delete from the end

However, there is no direct deletion of the object. In fact, we just change the links at the nodes, thereby delete items from the list.

There will be no code examples here, but there will be a deletion algorithm:

#### Delete from the start:

1. Move `head` to `head.next`
2. Thats all :D

#### Delete from the middle:

1. `dummy = head`
2. Traverse to desired element by stopping before it
3. So `dummy.next.val` must be desired element
4. `dummy.next = dummy.next.next`
5. Thats all :D

#### Delete from the end:

This is a special case of removal from the middle. The only difference is that after deleting, you need to change the link from `tail`.

1. `dummy = head`
2. Traverse to desired element by stopping before it
3. `if dummy.next == tail -> tail = dummy`
4. `dummy.next = dummy.next.next`
5. Thats all :D


### Searching

The search for an item in the list can be described by the following algorithm:

1. `dummy = head`
2. Traverse to `tail` of list
3. While traverse -> compare `dummy.val and item`
4. `if dummy,val == item -> return its position`
5. If the cycle is over, then there is no such element -> `return -1`

### Traversal

It may have other names, but the main purpose of such a function is to go through the entire list and output or return all its values in order.

Algorithm:

1. `dummy = head`
2. Traverse to `tail` of list
3. Print `dummy.val` OR save it to `array`
4. `Return array` of values (if you didn't print it)

### Size

Algorithm:

1. `dummy = head`, `counter = 0`
2. Traverse to `tail` of list, `counter++`
3. `Return counter + 1`

## Implementation

This implementation is written entirely in the Golang language without using third-party libraries.

_TODO: add threadsafe_

Interfaces and auxiliary structures are described in the file [list.go](/list/list.go). After going to it, you can read the comments near the method signatures to understand what they do

We are interested in the following structures and interfaces:
```go
// node is a struct of list to store val and ptr to next node
type node[T comparable] struct {
	val  T
	next *node[T]
}

// Common List interface with common operations
type List[T comparable] interface {
	Add(item T) error
	Insert(item T, pos int) error
	RemoveLast() error
	RemoveVal(item T) (int, error)
	RemoveAt(pos int) error
	Set(item T, pos int) error
	Get(pos int) (*T, error)
	GetLast() (*T, error)
	GetPosition(item T) (int, error)
	Size() int
	Clear() error
	Contains(item T) bool
}
```

You can see that generics are used. This is necessary so that the structure can store data of any type, however, unlike using `interface{}` as a data type, generics allow you to make typing strict. That is, you cannot store an `int` and a `string` in one instance of the structure. You can read more about generics in Golang here:

[Introduction to Generics in Golang](https://go.dev/doc/tutorial/generics)

It is recommended to use interfaces to declare variables and further initialize them using functions. We can find implementations of the **SLL** operations in this file: [singly_linked_list.go](/list/singly_linked_list.go)

It is worth noting that this implementation in its structure contains, in addition to head and tail, also size. This is necessary in order not to constantly go through the list and for some facilitating conditions.

## Usage

```go
package main

import (
	"fmt"

	"github.com/0x0FACED/go-collections/list"
)

func main() {
	var sll list.List[int]
	sll = list.NewSinglyLinked[int]()

	sll.Add(1)
	sll.Add(2)
	sll.Add(3)
	sll.Add(4)

	val, err := sll.Get(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(*val, " ") // Output: 1

	val, err = sll.Get(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(*val, " ") // Output: 3

	sll.Insert(15, 2)
	sll.Insert(50, 0)

	val, err = sll.Get(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(*val, " ") // Output: 50

	val, err = sll.Get(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*val, " ") // Output: 15
}

```

