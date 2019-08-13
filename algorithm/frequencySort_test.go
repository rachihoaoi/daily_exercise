package algorithm

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	fmt.Println(frequencySort("tree"))
}

func frequencySort(s string) string {
	m := make(map[int8]int)
	for i := 0; i < len(s); i++ {
		_, ok := m[int8(s[i])]
		if ok != true {
			m[int8(s[i])] = strings.Count(s, string(s[i]))
		}

	}
	//fmt.Println(m)
	n := map[int][]int8{}
	var a []int

	for k, v := range m {
		n[v] = append(n[v], k)
	}
	//fmt.Println(n)
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	res := ""
	for _, k := range a {
		for _, s := range n[k] {
			res += strings.Repeat(string(s), k)
		}
	}
	return res

}