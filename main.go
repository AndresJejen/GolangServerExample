package main

import (
	"fmt"

	"github.com/AndresJejen/AprendiendoGo/calculadora"
)

func main() {
	var saludo = "Hola estoy saludando"
	fmt.Printf("Hello World %v \n", saludo)

	var numero1 float32 = 12.3
	var numero2 float32 = 23

	fmt.Printf("Sumando %v y %v %v \n", numero1, numero2, calculadora.Sumar(numero1, numero2))
	fmt.Printf("Restando %v y %v %v \n", numero1, numero2, calculadora.Restar(numero1, numero2))
	fmt.Printf("Multiplicando %v y %v %v \n", numero1, numero2, calculadora.Multiplicar(numero1, numero2))

	var res, err = calculadora.Dividir(numero1, numero2)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Dividiendo %v y %v %v \n", numero1, numero2, res)
	}
}
