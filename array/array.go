package main

import (
	"fmt"
	"sort"
)

func main() {

	// Array
	students := [5]string{"Rafael", "Gustavo", "Gabriel", "Catarina", "Cibelle"}
	sort.Strings(students[:])
	for i := 0; i < len(students); i++ {
		fmt.Println("Aluno: ", students[i])
	}

	// Slices
	topTwoStudents := students[0:2]

	for i := 0; i < len(topTwoStudents); i++ {
		fmt.Printf("Student %v: %v\n", i+1, topTwoStudents[i])
	}

	// Slices com tamanhos dinâmicos
	dynamicSlice := make([]int, 10)

	for i := 0; i < len(dynamicSlice); i++ {
		fmt.Printf("Item %v: %v\n", i+1, dynamicSlice[i])
	}

	// Custom Array
	var myArray MyArray
	myArray.append(1)
	myArray.append(2)
	myArray.append(3)
	myArray.append(4)
	myArray.append(5)
	myArray.append(6)
	myArray.append(7)

	// String Array

	var strArray StrArray
	strArray.append("a")
	strArray.append("b")
	strArray.append("c")
	strArray.append("d")
	strArray.append("e")

	fmt.Printf("strArray: %v\n", strArray)

	strArray.reverse()
	fmt.Printf("strArray: %v\n", strArray)

}
