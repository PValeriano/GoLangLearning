package main

import "fmt"

func main() {
	valuesAndVariables()
	forLoop()
	condicionales(2)
	condicionales(3)
	mySwitch(3)
	mySwitch(127)
	myRange([3]int8{3, 0, 17})
}

func valuesAndVariables() {
	var text = "abc"
	fmt.Println("Strings: ", text)
	fmt.Println("Enteros: ", 400)
	fmt.Println("Decimales: ", 4.5)
	//Podemos causar un overflow si excedemos el tamaño permitido
	//var integer8 int8 = 127 + 1
	var integer8 int8 = 127
	fmt.Println("Int8: ", integer8)
	var ingeger16 int16 = 32767
	fmt.Println("Int16: ", ingeger16)
	fmt.Println("Booleanos: ", true, false)
	var myArray = [3]int8{1, 2, 3}
	fmt.Println("Array lineal: ", myArray)
	//Inicializamos vacío un array de 2x3
	var multidimensionalArray [2][3]int
	//Poblamos las posiciones vacías del array de una en una.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			multidimensionalArray[i][j] = i + j
		}
	}
	fmt.Println("Array bidimensional: ", multidimensionalArray)

	var mySlice = []int{4, 5, 6}
	fmt.Println("Slice: ", mySlice)
}

func forLoop() {
	fmt.Println("Forma 1: definir iterador + aumentar manualmente iterador.")
	var i = 1
	for i <= 3 {
		fmt.Println("Iteracion: ", i)
		i = i + 1
	}
	fmt.Println("Forma 2: utilizar iterador y aumentar directamente")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
}

func condicionales(a int8) {
	if a%2 == 0 {
		fmt.Println("Número par: ", a)
	} else {
		fmt.Println("Número impar.")
	}
}

func mySwitch(i int8) {
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("No considerado")
	}
}

func myRange(nums [3]int8) {
	for i, num := range nums {
		fmt.Println("Indice: ", i)
		fmt.Println("Número: ", num)
		if num == 3 {
			fmt.Println("Un tres! Qué suerte")
		}
	}
}
