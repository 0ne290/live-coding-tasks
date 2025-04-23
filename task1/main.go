package main

import "fmt"

// На вход подаются два неупорядоченных массива любой длины. Необходимо написать функцию, которая возвращает
// пересечение массивов.

func intersection(a, b []int) []int {
	lengthA := len(a)
	lengthB := len(b)
	var maxLength int
	if lengthA > lengthB {
		maxLength = lengthA
	} else {
		maxLength = lengthB
		a, b = b, a
	}

	ret := make([]int, 0, maxLength)
	difference := make(map[int]bool, maxLength)

	for _, item := range a {
		difference[item] = true
	}
	for _, item := range b {
		if difference[item] {
			ret = append(ret, item)
			delete(difference, item)
		}
	}

	return ret
}

func main() {
	a := []int{1, 2, 2, 3, 3, 4, 4, 5}
	b := []int{3, 3, 4, 4, 6, 7, 8, 9, 1}

	fmt.Println(intersection(a, b))
}
