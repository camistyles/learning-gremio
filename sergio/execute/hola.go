package main

import (
	"fmt"
	"github.com/learning-gremio/sergio/comon"
	"os"
)

func main() {
	var numero1 int
	fmt.Println("colcar el primer numero por favor\n")
	_, err := fmt.Scan(&numero1)
	if err != nil{
		panic(err)
	}

	var numero2 int
	fmt.Println("coloca el segundo numero\n")

	_, err = fmt.Scan(&numero2)

	err, result := comon.Sumar(numero1, numero2)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("el resultado de la suma es ", result)
	os.Exit(0)



	//fmt.Println(fmt.Sprintf("hola este es el numero 1 %v y este es el numero 2 %v", numero1, numero2))
}

