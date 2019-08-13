package algorithm

import (
	"fmt"
	"testing"
)


func TestIsValid(t *testing.T) {
	fmt.Println(isValid("([)]"))
}

func PushS(s *[]byte, v byte) {
	*s = append(*s, v)
}

func PopS(s *[]byte) byte {
	if len(*s) == 0 {
		return 0
	}
	len := len(*s)
	v := (*s)[len-1]
	*s = (*s)[:len-1]

	return v
}

func GetTop(s *[]byte) byte {
	len := len(*s)
	if len > 0 {return (*s)[len-1]}
	return 0

}


func isValid(s string) bool {
	array := []byte(s)
	fmt.Println(array)
	stack := make([]byte, 0)
	for i := 0; i < len(array); i++ {
		if array[i] - GetTop(&stack) <= 2 && array[i] - GetTop(&stack) != 0{
			PopS(&stack)
		} else {
			PushS(&stack, array[i])
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}