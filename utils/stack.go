package utils

import "errors"

// Stack type
type Stack[T interface{}] struct {
	elements []T
	n        int
}

// NewStack func
func NewStack[T interface{}](elements []T) *Stack[T] {
	st := Stack[T]{elements, len(elements)}
	return &st
}

// Push func
func (st *Stack[T]) Push(element T) {
	st.elements = append(st.elements, element)
	st.n++
}

// PushMulti func reserves the order of elements
func (st *Stack[T]) PushMulti(elements []T) {
	st.elements = append(st.elements, elements...)
	st.n += len(elements)
}

// Pop func
func (st *Stack[T]) Pop() (T, error) {
	element, err := st.Top()
	if err != nil {
		return element, err
	}
	st.elements = st.elements[:st.n-1]
	st.n--
	return element, nil
}

// PopMulti func
func (st *Stack[T]) PopMulti(len int) ([]T, error) {
	if st.n < len {
		return nil, errors.New("not enough element")
	}
	subElements := make([]T, len)
	copy(subElements, st.elements[st.n-len:])
	st.elements = st.elements[:st.n-len]
	st.n -= len
	return subElements, nil
}

// Top func
func (st *Stack[T]) Top() (T, error) {
	var element T
	if st.IsEmpty() {
		return element, errors.New("no element to get")
	}
	element = st.elements[st.n-1]
	return element, nil
}

// IsEmpty func
func (st *Stack[T]) IsEmpty() bool {
	return st.n == 0
}
