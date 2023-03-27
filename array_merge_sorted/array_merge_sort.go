package main

import (
	"fmt"
	"math"
)

func mergeStoredArrays(sortedArray1, sortedArray2 []int) []int {
	var i, j int = 1, 1
	// Números positivos?
	// Números de tipos inteiros e que não se repetem?
	// Checar as entradas
	if len(sortedArray1) == 0 {
		return sortedArray2
	}
	if len(sortedArray2) == 0 {
		return sortedArray1
	}
	lenArr1, lenArr2 := len(sortedArray1), len(sortedArray2)
	merged := make([]int, 0)
	currentItemAr1 := sortedArray1[0]
	currentItemAr2 := sortedArray2[0]
	hasItemsArr1, hasItemsArr2 := true, true

	for hasItemsArr1 || hasItemsArr2 {
		fmt.Printf("Valor 1: %v - Valor 2: %v\n", currentItemAr1, currentItemAr2)
		if currentItemAr1 <= currentItemAr2 {
			merged = append(merged, currentItemAr1)
			if i+1 <= lenArr1 {
				currentItemAr1 = sortedArray1[i]
				i += 1
			} else {
				hasItemsArr1 = false
				currentItemAr1 = math.MaxInt
			}
		} else {
			merged = append(merged, currentItemAr2)
			if j+1 <= lenArr2 {
				currentItemAr2 = sortedArray2[j]
				j += 1
			} else {
				hasItemsArr2 = false
				currentItemAr2 = math.MaxInt
			}
		}
	}
	return merged
}

func main() {
	sortedArray1 := [4]int{0, 3, 4, 31}
	sortedArray2 := [3]int{4, 6, 30}

	merge := mergeStoredArrays(sortedArray1[:], sortedArray2[:])

	for i := 0; i < len(merge); i++ {
		fmt.Printf("%v\n", merge[i])
	}
}
