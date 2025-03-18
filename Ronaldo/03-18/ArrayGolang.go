//! array - vetor - em Go são estáticos
//* ex1:
//? 	var arraydosnumeros [5]int
//* ex2:
//? 	var arraydosnumeros := [5] int{}
//* ex3:
//? 	var arraydosnumeros := [5] int{1,2,3,4,5}
// --------------------------------------------
//! Slice
//* ex1:
//? 	filmes := make([]string,3)
//! Propriedades Slice
//? 	Índice (ponteiro)
//?		len (length)
//? 	cap (capacidade)
// --------------------------------------------
//! Append
//* append(nome do Slice, valor a ser adicionado)
//? append(filmes, "filme1")
// --------------------------------------------
package main

import "fmt"

func main() {
	//* empy array
	myArray := [5]int{}
	fmt.Println(myArray)
	//* slice array
	mySlice := make([]int, 5)
	fmt.Println(mySlice)
	//* append
	mySlice = append(mySlice, 2)
	fmt.Println(mySlice)
	//* for
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("Indice: %v com valor %v \n", i, mySlice[i])
	}
}
