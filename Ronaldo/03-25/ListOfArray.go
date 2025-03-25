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
	//* short for
	for index, value := range mySlice {
		fmt.Printf("Indice: %v com valor %v \n", index, value)
	}
	//* map
	animals := map[string]string{
		"nome": "Rex",
		"tipo": "Cachorro",
		"cor": "Preto",
	}
	for char, avalue := range animals {
		fmt.Printf("chave: %v com valor %v \n", char, avalue)
	}
}
