package main

import "fmt"

func main() {
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
	fmt.Println("Array: ", myArray)
	var mySlice = []int{4, 5, 6}
	fmt.Println("Slice: ", mySlice)
}
