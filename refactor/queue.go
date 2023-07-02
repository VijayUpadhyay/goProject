package main

import (
	"sync/atomic"
	"unsafe"
)

// LockfreeQueue is a goroutine-safe Queue implementation.
// The overall performance of LockfreeQueue is much better than List+Mutex(standard package).
type LockfreeQueue[T any] struct {
	head  unsafe.Pointer
	tail  unsafe.Pointer
	dummy lfqNode[T]
}

// NewLockfreeQueue is the only way to get a new, ready-to-use LockfreeQueue.
//
// Example:
//
//	lfq := queue.NewLockfreeQueue[int]()
//	lfq.Push(100)
//	v, ok := lfq.Pop()
func NewLockfreeQueue[T any]() *LockfreeQueue[T] {
	var lfq LockfreeQueue[T]
	lfq.head = unsafe.Pointer(&lfq.dummy)
	lfq.tail = lfq.head
	return &lfq
}

// Pop returns (and removes) an element from the front of the queue and true if the queue is not empty,
// otherwise it returns a default value and false if the queue is empty.
// It performs about 100% better than list.List.Front() and list.List.Remove() with sync.Mutex.
func (lfq *LockfreeQueue[T]) Pop() (T, bool) {
	for {
		h := atomic.LoadPointer(&lfq.head)
		rh := (*lfqNode[T])(h)
		n := (*lfqNode[T])(atomic.LoadPointer(&rh.next))
		if n != nil {
			if atomic.CompareAndSwapPointer(&lfq.head, h, rh.next) {
				return n.val, true
			} else {
				continue
			}
		} else {
			var v T
			return v, false
		}
	}
}

// Push inserts an element to the back of the queue.
// It performs exactly the same as list.List.PushBack() with sync.Mutex.
func (lfq *LockfreeQueue[T]) Push(val T) {
	node := unsafe.Pointer(&lfqNode[T]{val: val})
	for {
		rt := (*lfqNode[T])(atomic.LoadPointer(&lfq.tail))
		//t := atomic.LoadPointer(&lfq.tail)
		//rt := (*lfqNode[T])(t)
		if atomic.CompareAndSwapPointer(&rt.next, nil, node) {
			atomic.StorePointer(&lfq.tail, node)
			// If dead loop occurs, use CompareAndSwapPointer instead of StorePointer
			// atomic.CompareAndSwapPointer(&lfq.tail, t, node)
			return
		} else {
			continue
		}
	}
}

type lfqNode[T any] struct {
	val  T
	next unsafe.Pointer
}
