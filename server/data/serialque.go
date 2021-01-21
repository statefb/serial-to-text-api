package data

import "errors"

type que []string

func (q que) len() int { return len(q) }
func (q *que) push(x interface{}) {
	*q = append(*q, x.(string))
}
func (q *que) pop() (interface{}, error) {
	old := *q
	n := len(old)
	if n == 0 {
		return nil, errors.New("que is empty")
	}
	x := old[n-1]
	*q = old[0 : n-1]
	return x, nil
}
