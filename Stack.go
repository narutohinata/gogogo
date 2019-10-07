package main

import "fmt"

type Stack struct {
	size uint
	capacity uint
	elements []int
}

func NewStack(capacity uint)*Stack  {
	return &Stack{
		size:     0,
		capacity: capacity,
		elements: make([]int, capacity),
	}
}

func (stack *Stack) StackEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) StackSize() uint  {
	return stack.size
}

func (stack *Stack) StackFull() bool  {
	return stack.size == stack.capacity
}

func (stack *Stack) Push(data int)  {
	if stack.StackFull() {
		stack.capacity *= 2
		newElements := make([]int, stack.capacity)
		for index, value := range stack.elements {
			newElements[index] = value
		}
		stack.elements = newElements
	}
	stack.size++
	stack.elements[stack.size-1] = data
}

func (stack *Stack) Pop() int {
	stack.size--
	value := stack.elements[0]
	for i := uint(0); i < stack.size ; i++  {
		stack.elements[i] = stack.elements[i+1]
	}
	return value
}

func (stack *Stack) Top() int  {
	return stack.elements[stack.size-1]
}

func (stack *Stack) PrintStack()  {
	for index := stack.size; index > 0; index-- {
		fmt.Printf("stack[%d] is %d\n", index, stack.elements[index-1])
	}
}