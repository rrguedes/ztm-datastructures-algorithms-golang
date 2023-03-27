package main

import "fmt"

type MyArray struct {
	length int
	data   []int
}

func (ar *MyArray) get(index int) int {
	return ar.data[index]
}

func (ar *MyArray) append(item int) {
	ar.data = append(ar.data, item)
	ar.length = len(ar.data)
}

func (ar *MyArray) print() {
	for i := 0; i < len(ar.data); i++ {
		fmt.Printf("%v", ar.get(i))
	}
}

type StrArray struct {
	length int
	data   []string
}

func (ar *StrArray) append(item string) {
	ar.data = append(ar.data, item)
	ar.length = len(ar.data)
}

func (ar *StrArray) get(index int) string {
	return ar.data[index]
}

func (ar *StrArray) reverse() []string {
	lastIndex := ar.length - 1
	reversedArray := make([]string, ar.length)

	for i := 0; i < ar.length; i++ {
		currentItem := ar.get(i)
		reversedArray[lastIndex] = currentItem
		lastIndex -= 1
	}
	ar.data = reversedArray
	return ar.data
}
