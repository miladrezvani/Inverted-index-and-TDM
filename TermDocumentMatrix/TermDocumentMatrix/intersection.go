package termdocumentmatrix

import (
	"slices"
	"strconv"
)

func intersection(hashmap map[string]map[string]int,data map[string][]int,keys int,val1 string , val2 string) []int {
	var intersection []int
	i := 0
	for i <= keys {
		if hashmap[strconv.Itoa(i)][val1] == 1 && hashmap[strconv.Itoa(i)][val2] == 1 {
			intersection = append(intersection, i)
		} else if (slices.Contains(data[val1], i) && hashmap[strconv.Itoa(i)][val2] == 1) || (hashmap[strconv.Itoa(i)][val1] == 1 && slices.Contains(data[val2], i)) {
			intersection = append(intersection, i)
		} 
		i++
	}

	return intersection
 }


func intersection_not(hashmap map[string]map[string]int,data map[string][]int,keys int,val1 string , val2 string, side bool) []int {
	var intersection []int
	i := 0
	if side {
		for i <= keys {
			if hashmap[strconv.Itoa(i)][val1] == 1 && hashmap[strconv.Itoa(i)][val2] == 1 {
				i++
				continue
			} else if hashmap[strconv.Itoa(i)][val1] == 1 && hashmap[strconv.Itoa(i)][val2] == 0 {
				intersection = append(intersection, i)
			} else if (slices.Contains(data[val1], i) && hashmap[strconv.Itoa(i)][val2] == 0) || (hashmap[strconv.Itoa(i)][val1] == 1 && !slices.Contains(data[val2], i)) {
				intersection = append(intersection, i)
			} 
			i++
		}
	}else {
		for i <= keys {
			if hashmap[strconv.Itoa(i)][val1] == 1 && hashmap[strconv.Itoa(i)][val2] == 1 {
				i++
				continue
			} else if hashmap[strconv.Itoa(i)][val1] == 0 && hashmap[strconv.Itoa(i)][val2] == 1 {
				intersection = append(intersection, i)
			} else if (!slices.Contains(data[val1], i) && hashmap[strconv.Itoa(i)][val2] == 1) || (hashmap[strconv.Itoa(i)][val1] == 0 && slices.Contains(data[val2], i)) {
				intersection = append(intersection, i)
			} 
			i++
		}
	}

	return intersection
 }

