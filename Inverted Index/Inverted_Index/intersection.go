package Inverted_Index

import (
	"slices"
	"strconv"
)

func intersection(arr1 []string , arr2 []string) []string {
	var intersection []string
	i := 0
	j := 0
	var A []string
	var B []string
	if len(arr1) < len(arr2) {
		A = arr1
		B = arr2
	}else {
		A = arr2
		B = arr1
	}
	for i < len(A) && j < len(B)  {
		first , _ := strconv.Atoi(A[i])
		second , _ := strconv.Atoi(B[j])
		if first < second {
			i++
		}else if first > second {
			j++
		}else {
			intersection = append(intersection, A[i])
			i++
			j++
		}
	}
	return intersection
 }


func intersection_not(arr1 []string , arr2 []string, side bool) []string {
	var intersection []string
	i := 0
	j := 0
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
		for i < len(A) && j < len(B)  {
			first , _ := strconv.Atoi(A[i])
			second , _ := strconv.Atoi(B[j])
			if first < second {
				if !slices.Contains(B, A[i]) && !slices.Contains(intersection, A[i]) {
					intersection = append(intersection, A[i])
				}
				i++
			}else if first > second {
				if !slices.Contains(B, A[i]) && !slices.Contains(intersection, A[i]) {
					intersection = append(intersection, A[i])
				}
				j++
			}else {
				i++
				j++
			}
		}
	}else {
		for i < len(A) && j < len(B)  {
			first , _ := strconv.Atoi(A[i])
			second , _ := strconv.Atoi(B[j])
			if first < second {
				if !slices.Contains(A, B[j]) && !slices.Contains(intersection, B[j]) {
					intersection = append(intersection, B[j])
				}
				i++
			}else if first > second {
				if !slices.Contains(A, B[j]) && !slices.Contains(intersection, B[j]) {
					intersection = append(intersection, B[j])
				}
				j++
			}else {
				i++
				j++
			}
		}
	}

	return intersection
 }

