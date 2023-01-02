package slicesmanager_test

import (
	"math/rand"
	"time"
)

func newRandGenerator() *rand.Rand {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r
}

func NewOrderedSlice(slice_size int) []int {

	s := make([]int, slice_size, slice_size)

	for i := 0; i < slice_size; i++ {
		s[i] = i
	}

	return s
}

func NewUnorderedSlice(slice_size int) []int {

	r := newRandGenerator()

	unordered_slice := make([]int, 0, slice_size)
	ordered_boxes := NewOrderedSlice(slice_size)

	for i := slice_size; i > 0; i-- {
		pos := r.Intn(i)
		num := ordered_boxes[pos]
		unordered_slice = append(unordered_slice, num)
		ordered_boxes = append(ordered_boxes[:pos], ordered_boxes[pos+1:]...)
	}

	return unordered_slice

}
