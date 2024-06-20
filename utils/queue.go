package utils

type Queue[T any] struct {
	elems []*T
}

func (q *Queue[T]) Peek() *T {
	if len(q.elems) == 0 {
		return nil
	}

	return q.elems[0]
}

func (q *Queue[T]) Add(elem *T) {
	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) Remove() *T {
	if len(q.elems) == 0 {
		return nil
	}

	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue[T]) Size() int {
	return len(q.elems)
}
