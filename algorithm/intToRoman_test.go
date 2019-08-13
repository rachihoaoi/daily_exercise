package algorithm

import (
	"fmt"
	"testing"
)

func TestIntTORoman(t *testing.T) {
	fmt.Println(intToRoman(123))
}

func intToRoman(num int) string {
	KeyMap := []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	keyValue := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	result := ""
	i := 0
	for num > 0 {
		if num < keyValue[i] {
			i++
		} else {
			result += KeyMap[i]
			num -= keyValue[i]
		}
	}
	return  result
}
