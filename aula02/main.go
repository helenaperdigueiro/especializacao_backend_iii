package main

import "fmt"

//guiado
const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	//guiado
	minFunc, err := operation(minimum)
	if err != nil {
		panic(err)
	}

	averageFunc, err := operation(average)
	if err != nil {
		panic(err)
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Valor minimo:", minValue)
	fmt.Println("Valor medio:", averageValue)
	fmt.Println("Valor maximo:", maxValue)

	//mesa
	salarioEmpregadoA, _ := calcularSalario(600, "A")
	salarioEmpregadoB, _ := calcularSalario(600, "B")
	salarioEmpregadoC, _ := calcularSalario(600, "C")
	test, err := calcularSalario(600, "D")
	if err != nil {
		panic(err)
	}

	fmt.Println("Salario do empregado A:", salarioEmpregadoA)
	fmt.Println("Salario do empregado B:", salarioEmpregadoB)
	fmt.Println("Salario do empregado C:", salarioEmpregadoC)
	fmt.Println("Salario do empregado C:", test)

}

//guiado
func operation(op string) (func(i ...int) int, error) {
	switch op {
	case minimum:
		return minCalc, nil
	case average:
		return averageCalc, nil
	case maximum:
		return maxCalc, nil
	}

	return nil, fmt.Errorf("operation not allowed")
}

func minCalc(i ...int) int {
	min := i[0]
	for _, v := range i {
		if (v < min) {
			min = v
		}
	}
	return min
}

func averageCalc(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum/len(i)
}

func maxCalc(i ...int) int {
	max := i[0]
	for _, v := range i {
		if max < v {
			max = v
		}
	}
	return max
}

//mesa
func calcularSalario(minutos float32, categoria string) (float32, error) {
	const (
		A = 112.38
		B = 56.19
		C = 37.46
	)
	var horas float32 = minutos/60

	switch categoria {
	case "A":
		return multiplicar(horas, A, 0.5), nil
	case "B":
		return multiplicar(horas, B, 0.2), nil
	case "C":
		return multiplicar(horas, C, 0), nil
	}
	return 0.0, fmt.Errorf("Erro no calculo do salario. Verifique os parametros inseridos.")
}

func multiplicar(horas float32, valorPorHora float32, porcentagem float32) float32 {
	return horas * valorPorHora * (1 + porcentagem)
}