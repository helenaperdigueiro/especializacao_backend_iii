package main

import "fmt"

func main() {
	//exercicio 1
	// word := "palavra"
	// fmt.Println(word)
	// fmt.Println("A palavra tem", len(word), "letras")

	// for count := 0; count < len(word); count ++ {
	// 	fmt.Println(string(word[count]))
	// }

	//exercicio 2
	// number := 5
	// months := []string{"Janeiro", "Fevereiro", "Marco", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	// fmt.Println(months[number-1])

	// number := 5
	// months := map[int]string{
	// 	1: "Janeiro", 
	// 	2: "Fevereiro", 
	// 	3: "Marco", 
	// 	4: "Abril", 
	// 	5: "Maio", 
	// 	6: "Junho", 
	// 	7: "Julho", 
	// 	8: "Agosto", 
	// 	9: "Setembro", 
	// 	10: "Outubro", 
	// 	11: "Novembro", 
	// 	12:"Dezembro"}
	// fmt.Println(months[number])

	//mesa
	//exercicio 1
	estudantes := []string{"Benjamin", "Fernando", "Brenda", "Marcos", "Pedro", "Evandro", "Alex", "Maria", "Frederico", "Juan", "Leandro", "Eduardo", "Camila"}
	estudantes = append(estudantes, "Gabriela")
	fmt.Println(estudantes)

	//exercicio 2
	var empregados = map[string]int{"Benjamin": 20, "Fernando": 28, "Brenda": 19, "Marcos": 44, "Pedro": 30}
	fmt.Println("A idade do Benjamin é", empregados["Benjamin"])

	var count int
	for _, idade := range empregados {
		if idade > 21 {
			count++
		}
	}
	fmt.Println(count, "funcionários têm 21 anos")

	empregados["Igor"] = 25
	fmt.Println(empregados)

	delete(empregados, "Pedro")
	fmt.Println(empregados)

}