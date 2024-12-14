package calc

import f "fmt"

func Calc() {
	var num1, num2 int
	f.Println("Bem vindo(a) a calculadora do tutuixa")

	f.Println("Digite o primeiro número")
	f.Scanln(&num1)

	f.Println("Digite o segundo número")
	f.Scanln(&num2)

	var sum int = num1 + num2
	var sub int = num1 - num2

	f.Println("Resultado")
	f.Println("Soma:", sum)
	f.Println("Subtração:", sub)
}
