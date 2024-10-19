package Inverted_Index

import (
	"slices"
	"strconv"
)

 func union(arr1 []string, arr2 []string) []string {

    i := 0
    j := 0
    var union []int
	var data []string
	var A []string
	var B []string
	if len(arr1) < len(arr2) {
		A = arr1
		B = arr2
	}else {
		A = arr2
		B = arr1
	}
    for i < len(A) && j < len(B) {
		first , _ := strconv.Atoi(A[i])
		second , _ := strconv.Atoi(B[j])
        if first <= second{
			if len(union) == 0 || union[len(union) - 1] != first {
				union = append(union, first)
				data = append(data, A[i])
			}
			i++
		}else {
            if len(union) == 0 || union[len(union) - 1] != second {
				union = append(union, second)
				data = append(data, B[j])
			}
			j++
		}
	}
    for i < len(A){
		first , _ := strconv.Atoi(A[i])
        if union[len(union) - 1] != first{
			union = append(union, first)
			data = append(data, A[i])
		}
        i++
	}

    for j < len(B) {
		second , _ := strconv.Atoi(B[j])
        if union[len(union) - 1] != second {
			union = append(union, second)
			data = append(data, B[j])
		}
        j++
	}
    return data
}


func union_not(arr1 []string, arr2 []string,side bool) []string {

    i := 0
	var data []string
	var A []string
	var B []string
	if len(arr1) < len(arr2) {
		A = arr1
		B = arr2
	}else {
		A = arr2
		B = arr1
	}
	if side {
		for i < 100000 {
			if !slices.Contains(B, strconv.Itoa(i)) {
				data = append(data, strconv.Itoa(i))
			}
			i++
		}
	}else {
		for i < 100000 {
			if !slices.Contains(A, strconv.Itoa(i)) {
				data = append(data, strconv.Itoa(i))
			}
			i++
		}
	}

    return data
}
