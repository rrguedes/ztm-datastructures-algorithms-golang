package main

import (
	"fmt"
	"math"
)

func mergeStoredArrays(sortedArray1, sortedArray2 []int) []int {
	var i, j int = 1, 1
	// Crio um slice vazio para ir inserindo os dados
	merged := make([]int, 0)
	// Para não ficar repetindo os lens, eu tomo o tamanho das arrays no início
	lenArr1, lenArr2 := len(sortedArray1), len(sortedArray2)

	// A primeira coisa que eu preciso fazer é tratar dos casos onde uma (ou ambas as) entradas seja(m) vazia(s). Neste caso eu retorno imediatamente a outra entrada ou o slice vazio
	if lenArr1 == 0 && lenArr2 == 0 {
		return merged //slice vazio
	}
	if lenArr1 == 0 {
		return sortedArray2
	}
	if lenArr2 == 0 {
		return sortedArray1
	}
	// Tomo os itens iniciais de cada array
	currentItemAr1 := sortedArray1[0]
	currentItemAr2 := sortedArray2[0]
	// Controlarei a condição utilizando variáveis que indicam se uma das arrays ainda possuem items
	hasItemsArr1, hasItemsArr2 := true, true

	// Se tiverem
	for hasItemsArr1 || hasItemsArr2 {
		fmt.Printf("Valor 1: %v - Valor 2: %v\n", currentItemAr1, currentItemAr2) // Log para validar o funcionamento - importante ir validando sua solução passo a passo durante a sua construção
		// Verifico se o item do primeiro array é menor ou igual do que a da segunda (visto que podem se repetir)
		if currentItemAr1 <= currentItemAr2 {
			// Caso o item 1 seja menor, incluo o mesmo no array
			merged = append(merged, currentItemAr1)
			// Imediatamente verifico se ele é o último e se não for, atualizo o item atual da Array 1, incrementando o contador
			if i+1 <= lenArr1 {
				currentItemAr1 = sortedArray1[i]
				i += 1
			} else {
				// Caso contrário, indico que não há próximo item. Como a array 2 ainda poderá ter um item, para que a comparação não falhe, atribuo o item atual da array 1 com o maior número possível
				hasItemsArr1 = false
				currentItemAr1 = math.MaxInt
			}
			// Caso contrário, o item atual do segundo array é menor, então faço o mesmo processo, só que para o item do array 2
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
	// Retorno a array ordenada
	return merged
}

func main() {
	/* O entrevistador pede para que você escreva uma função que mescle duas arrays ordenadas diferentes */
	sortedArray1 := [4]int{0, 3, 4, 31}
	sortedArray2 := [3]int{4, 6, 30}

	/* 	Na realidade o entrevistador quer entender o quão cuidadoso você é, por isso é importante validar
	alguns pontos antes de efetivamente escrever o seu código, como:
	-	Pergunte sobre o tipo de dados das arrays?
	-	Os números serão sempre positivos, negativos?
	-	Os números podem ser repetir?
	...

	Com estes questionamentos básicos validados é importante que você faça com que o entrevistador
	acompanhe a sua linha de raciocínio. Então não se acanhe e narre o seu desenvolvimento
	visto que saber lidar com situações de tensão com calma, garantindo que você consegue se
	comunicar, é um dos aspectos importantes na entrevista.

	Geralmente durante uma entrevista não conseguimos implementar a melhor solução, mas é importante deixar
	o entrevistador saber que você sabe o que fazer para tornar o seu código melhor. No caso em questão isso
	poderia ser extrair funções da função mergeStoredArrays tornando suas funções menores e mais legíveis

	*/

	merge := mergeStoredArrays(sortedArray1[:], sortedArray2[:])

	for i := 0; i < len(merge); i++ {
		fmt.Printf("%v\n", merge[i])
	}
}
