package algorithm

import (
	"fmt"
	"testing"
)

func TestDTR (t *testing.T) {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	//KeyMap := []string{"M","D","C","L","X","V","I"}
	//keyValue := []int{1000,500,100,50,10,5,1}
	hMap := map[string]int {
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	sum := 0
	for i := 0; i< len(s) - 1; i++ {
		if hMap[string(s[i])] >= hMap[string(s[i + 1])] {
			sum += hMap[string(s[i])]
		}else {
			sum -= hMap[string(s[i])]
		}
	}
	sum += hMap[string(s[len(s) - 1])]
	return sum
}
