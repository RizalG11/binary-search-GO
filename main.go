package main

import "fmt"

type RecType struct {
	count int
	size  int
}
type ArrType [100]RecType

func main() {
	var found int
	var min, mid, max int
	max = 100
	found = -1
	min = 0

	for min <= max && found == -1 {
		mid = (min + max) / 2
		if v > tab[mid].count {
			min = mid + 1
		} else if v < tab[mid].count {
			max = mid - 1
		} else {
			found = mid
		}
	}
	if found != -1 {
		return true
	} else {
		return false
	}

}
