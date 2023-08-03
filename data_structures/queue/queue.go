package queue

type Queue struct {
	arr []int
}

func (q *Queue) Enqueue(element int) {
	q.arr = append(q.arr, element)
}

func (q *Queue) Dequeue() int {
	val := q.arr[0]
	if len(q.arr) == 1 {
		q.arr = []int{}
	} else {
		q.arr = q.arr[1:]
	}
	return val
}
