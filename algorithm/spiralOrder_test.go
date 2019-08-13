package algorithm

import (
	"fmt"
	"testing"
	"time"
)

type Slice []int

func NewSlice() Slice {
	a := new(Slice)
	return *a
}

func (s *Slice) Add(e int) *Slice {
	*s = append(*s, e)
	fmt.Println(e)
	return s
}

func TestSpiralOrder (t *testing.T) {
	strs := []string{"one", "two", "three"}

	for _, s := range strs {
		go func(s string) {
			time.Sleep(1 * time.Second)
			fmt.Println(s)
		}(s)
	}

	time.Sleep(3 * time.Second)
}
