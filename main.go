// Шиляев Павел Юрьевич
package main

import "fmt"

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}

	uniqueA := unique(a)
	uniqueB := unique(b)
	intersection := intersect(uniqueA, uniqueB)
	union := merge(uniqueA, uniqueB)

	fmt.Println(uniqueA, uniqueB)
	fmt.Println(intersection)
	fmt.Println(union)
}

func unique(slice []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func intersect(a, b []int) []int {
	set := make(map[int]bool, len(a))
	for _, v := range a {
		set[v] = true
	}
	
	result := make([]int, 0)
	for _, v := range b {
		if set[v] {
			result = append(result, v)
			delete(set, v)
		}
	}
	return result
}

func merge(a, b []int) []int {
	seen := make(map[int]bool, len(a)+len(b))
	result := make([]int, 0, len(a)+len(b))
	
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}