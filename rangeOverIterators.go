package main

import "iter"

type List[T any] struct {
	head, tail *element[T]
}

//Head points to the firsd element of the list
//Tail points to the last element of the list

type element[T any] struct {
	next *element[T]
	val  T
}

//Element[T] is a node in the list
// val  is value of the node
// next is a pointer to the next node

// Adding Elements to the list
func (lst *List[T]) Push(v T) {
	//If the list is empty then create a new element and assign it to head and tail
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Iterating over the list
func (lst *List[T]) All() iter.Seq[T]{
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				break
			}
		}
	}
}