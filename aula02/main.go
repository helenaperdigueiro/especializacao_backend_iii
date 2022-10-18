package main

import "fmt"

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

	return nil, fmt.Errorf("operetion not allowed")
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
