package main

import (
	"fmt"
)

// Criar uma array com tamanho específico como parte de uma struct

type HashEntry struct {
	key   string
	value int
}

type HashTable struct {
	size int
	data []int
}

func (hash *HashTable) Size() int {
	return len(hash.data)
}

func (hash *HashTable) Set(key string, value int) {
	address := hash.hash(key)
	if hash.data[address] ==  {
		hash.data = append(hash.data[:address], HashEntry{key, value})
	}
}

func (hash *HashTable) Get(key string) HashEntry {
	address := hash.hash(key)
	return hash.data[address]
}

func (hash *HashTable) hash(key string) int {
	var hashKey int = 0
	runeKey := []rune(key)

	for i := 0; i < len(runeKey); i++ {
		hashKey = (hashKey + int(runeKey[i])*i) % (len(hash.data) + 1)
	}
	return hashKey
}

func main() {

	Prova := HashTable{}
	Prova.Set("Gustavo", 10)
	Prova.Set("José", 3)
	Prova.Set("Gabriel", 10)
	Prova.Set("Ricardo", 5)

	fmt.Println("Número de Alunos que realizaram a Prova: ", Prova.Size())

	notaGustavo := Prova.Get("Gustavo").value
	fmt.Println("Nota do Aluno Gustavo ", notaGustavo)
}
