// Aluno: Felipe Adrian Lourenço Barbisan
// Curso: 4º AMS (ADS) Fatec

package main

import (
	"fmt"
)

func main() {
	// EX 01 - Estudante
	// Não entendi o que era pra fazer

	// Ex 02 - Carro
	// 		c1 := Car{}
	// 		Registercar(&c1)
	// 		c2 := Car{}
	// 		Registercar(&c2)
	// 		c3 := Car{}
	// 		Registercar(&c3)
	// 		c4 := Car{}
	// 		Registercar(&c4)
	// 		c5 := Car{}
	// 		Registercar(&c5)
	// 		fmt.Println("Carro 1: ", c1)
	// 		fmt.Println("Carro 2: ", c2)
	// 		fmt.Println("Carro 3: ", c3)
	// 		fmt.Println("Carro 4: ", c4)
	// 		fmt.Println("Carro 5: ", c5)
	// 		ComsuptionCalculation(c1.Consumption);
	// 		ComsuptionCalculation(c2.Consumption);
	// 		ComsuptionCalculation(c3.Consumption);
	// 		ComsuptionCalculation(c4.Consumption);
	// 		ComsuptionCalculation(c5.Consumption);

	// }

	// func ComsuptionCalculation(c float32) {
	// 	fmt.Println("Digite o Km da viagem Pretendida:")
	// 	var km float32
	// 	fmt.Scanf("%s\n", &km)
	// 	fmt.Println("Digite o Preço do Combustivel:")
	// 	var fuel float32
	// 	fmt.Scanf("%s\n", &fuel)
	// 	calc := (km / c) * fuel
	// 	fmt.Println("O Gasto dessa viagem será de : R$", calc)
	// }
	// func Registercar(c *Car) {

	// 	fmt.Println("Digite o nome da marca do carro:")
	// 	fmt.Scanf("%s\n", &c.Brand)
	// 	fmt.Println("Digite o modelo do carro:")
	// 	fmt.Scanf("%s\n", &c.Model)
	// 	fmt.Println("Digite o ano do carro:")
	// 	fmt.Scanf("%d\n", &c.Year)
	// 	fmt.Println("Digite a identificação do carro:")
	// 	fmt.Scanf("%s\n", &c.Indetify)
	// 	fmt.Println("Digite o consumo do carro:")
	// 	fmt.Scanf("%f\n", &c.Consumption)
	// }
	// type Car struct{
	// 	Brand string
	// 	Model string
	// 	Year int
	// 	Indetify string
	// 	Consumption float32
	// }


	// Ex 03 - Funcionário
// 	e1 := Employee{}
// 	RegisterEmployee(&e1)
// 	e2 := Employee{}
// 	RegisterEmployee(&e2)
// 	e3 := Employee{}
// 	RegisterEmployee(&e3)
// 	e4 := Employee{}
// 	RegisterEmployee(&e4)
// 	e5 := Employee{}
// 	RegisterEmployee(&e5)
// 	e6 := Employee{}
// 	RegisterEmployee(&e6)
// 	e7 := Employee{}
// 	RegisterEmployee(&e7)
// 	e8 := Employee{}
// 	RegisterEmployee(&e8)
// 	e9 := Employee{}
// 	RegisterEmployee(&e9)
// 	e10 := Employee{}
// 	RegisterEmployee(&e10)
// 	ShowEmployee(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10)
// 	wageAverage(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10)
// }

// func wageAverage(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 Employee) {
// 	average := (e1.wage + e2.wage + e3.wage + e4.wage + e5.wage + e6.wage + e7.wage + e8.wage + e9.wage + e10.wage) / 10
// 	fmt.Println("A média salarial é: ", average)
// }

// func ShowEmployee(e1, e2, e3, e4, e5, e6, e7, e8, e9, e10 Employee) {

// 	fmt.Println("Funcionário 1: ", e1)
// 	fmt.Println("Salario Liquido:", e1.wage+e1.benefits-e1.discounts)
// 	fmt.Println("Funcionário 2: ", e2)
// 	fmt.Println("Salario Liquido:", e2.wage+e2.benefits-e2.discounts)
// 	fmt.Println("Funcionário 3: ", e3)
// 	fmt.Println("Salario Liquido:", e3.wage+e3.benefits-e3.discounts)
// 	fmt.Println("Funcionário 4: ", e4)
// 	fmt.Println("Salario Liquido:", e4.wage+e4.benefits-e4.discounts)
// 	fmt.Println("Funcionário 5: ", e5)
// 	fmt.Println("Salario Liquido:", e5.wage+e5.benefits-e5.discounts)
// 	fmt.Println("Funcionário 6: ", e6)
// 	fmt.Println("Salario Liquido:", e6.wage+e6.benefits-e6.discounts)
// 	fmt.Println("Funcionário 7: ", e7)
// 	fmt.Println("Salario Liquido:", e7.wage+e7.benefits-e7.discounts)
// 	fmt.Println("Funcionário 8: ", e8)
// 	fmt.Println("Salario Liquido:", e8.wage+e8.benefits-e8.discounts)
// 	fmt.Println("Funcionário 9: ", e9)
// 	fmt.Println("Salario Liquido:", e9.wage+e9.benefits-e9.discounts)
// 	fmt.Println("Funcionário 10: ", e10)
// 	fmt.Println("Salario Liquido:", e10.wage+e10.benefits-e10.discounts)
	
// }

// func RegisterEmployee(employee *Employee) {
// 	fmt.Println("Digite o nome do funcionário:")
// 	fmt.Scanf("%s\n", &employee.name)
// 	fmt.Println("Digite o cargo do funcionário:")
// 	fmt.Scanf("%s\n", &employee.position)
// 	fmt.Println("Digite o salário do funcionário:")
// 	fmt.Scanf("%f\n", &employee.wage)
// 	fmt.Println("Digite os benefícios do funcionário:")
// 	fmt.Scanf("%f\n", &employee.benefits)
// 	fmt.Println("Digite os descontos do funcionário:")
// 	fmt.Scanf("%f\n", &employee.discounts)
// }

// type Employee struct {
// 	name      string
// 	position  string
// 	wage      float32
// 	benefits  float32
// 	discounts float32
// }

// Ex 04 - Livro

//  	var livros []livro
// 	l1 := livro{
// 		titulo: "banana",
// 		autor: "J.R.R. Tolkien",
// 		ano: 1954,
// 		status: true,
// 	}
	
// 	l1.cadastrarLivro(&livros,l1) 
// 	l1.mostrarLivro(&livros)
// 	fmt.Println("Digite o título do livro:")
// 	var titulo string
// 	fmt.Scanf("%v\n", &titulo)
// 	pesquisarLivro(titulo, livros)
// 	l1.mudarStatus(&livros, l1)
// 	l1.mostrarLivro(&livros)

// }
// type livro struct {
// 	titulo string
// 	autor string
// 	ano int 
// 	status bool
// }
// func (l *livro) cadastrarLivro(livros *[]livro, livro livro) {
// 	*livros = append(*livros, livro)
// }
// func pesquisarLivro(t string, livros []livro) {
// 	for _, livro := range livros {
// 		if livro.titulo == t {
// 			fmt.Printf("Livro encontrado: %s, Autor: %s, Ano: %d, Status: %t\n", livro.titulo, livro.autor, livro.ano, livro.status)
// 			return
// 		}
// 	}
// 	fmt.Println("Livro não encontrado.")

// }
// func (l *livro)mostrarLivro(livros *[]livro){
// 	for i, livro := range *livros{
// 		fmt.Printf("Livro %d: %s, Autor: %s, Ano: %d, Status: %t\n", i+1, livro.titulo, livro.autor, livro.ano, livro.status)
// 	} 
// }
// func (l *livro) mudarStatus(livros *[]livro, livro livro) {
// 	for i, l := range *livros {
// 		if l.titulo == livro.titulo {
// 			(*livros)[i].status = !(*livros)[i].status
// 			fmt.Printf("Status do livro %s alterado para %t\n", livro.titulo, (*livros)[i].status)
// 			return
// 		}
// 	}
// 	fmt.Println("Livro não encontrado.")

// }