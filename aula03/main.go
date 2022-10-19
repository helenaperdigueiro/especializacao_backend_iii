package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	tarantula = "tarantula"
	hamster   = "hamster"
)


func main() {
	animalDog, error := animal(dog)
	if error != nil {
		panic(error)
	}
	amountDogFood := animalDog(2)
	fmt.Println("Quantidade de comida de cachorro:", amountDogFood, "g")

	animalCat, error := animal(cat)
	if error != nil {
		panic(error)
	}
	amountCatFood := animalCat(3)
	fmt.Println("Quantidade de comida de gato:", amountCatFood, "g")

	animalTarantula, error := animal(tarantula)
	if error != nil {
		panic(error)
	}
	amountTarantulaFood := animalTarantula(4)
	fmt.Println("Quantidade de comida de tarantula:", amountTarantulaFood, "g")

	animalHamster, error := animal(hamster)
	if error != nil {
		panic(error)
	}
	amountHamsterFood := animalHamster(5)
	fmt.Println("Quantidade de comida de hamster:", amountHamsterFood, "g")

	animalTest, error := animal("test")
	if error != nil {
		panic(error)
	}
	amountTest := animalTest(5)
	fmt.Println(amountTest)
}

func animal(animal string) (func(quantidadeAnimais int) int, error) {

	switch animal {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	case hamster:
		return hamsterFunc, nil
	}

	return nil, fmt.Errorf("Animal não encontrado")
}

func dogFunc(quantidadeAnimais int) int {
	return 10000 * quantidadeAnimais
}

func catFunc(quantidadeAnimais int) int {
	return 5000 * quantidadeAnimais
}

func tarantulaFunc(quantidadeAnimais int) int {
	return 150 * quantidadeAnimais
}

func hamsterFunc(quantidadeAnimais int) int {
	return 250 * quantidadeAnimais
}