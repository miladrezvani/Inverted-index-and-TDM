package termdocumentmatrix

import (
	"slices"
	"strconv"
)

 func union(hashmap map[string]map[string]int,data map[string][]int, keys int,val1 string, val2 string) []int {

    i := 0
    var union []int
	for i <= keys {
		if hashmap[strconv.Itoa(i)][val1] == 1 || hashmap[strconv.Itoa(i)][val2] == 1 || slices.Contains(data[val1], i) || slices.Contains(data[val2], i) {
			union = append(union, i)
		} 
		i++
	}
    return union
}


func union_not(hashmap map[string]map[string]int,data map[string][]int, keys int,val1 string, val2 string,side bool) []int {

    i := 0
	var union []int
	if side {
		for i <= keys {
			if hashmap[strconv.Itoa(i)][val2] == 0 || hashmap[strconv.Itoa(i)][val1] == 1 || slices.Contains(data[val1], i) {
				union = append(union, i)
			} 
			i++
		}
	}else {
		for i <= keys {
			if hashmap[strconv.Itoa(i)][val1] == 0 || hashmap[strconv.Itoa(i)][val2] == 1 || slices.Contains(data[val2], i) {
				union = append(union, i)
			} 
			i++
		}
	}

    return union
}
