package queue

type Queue []int

func (q *Queue) Push(num int) {
	for len(*q) > 0 && num >= (*q)[len(*q)-1] {
		*q = (*q)[:len(*q)-1]
	}
	*q = append(*q, num)
}
