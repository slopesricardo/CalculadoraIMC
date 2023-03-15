package main

import "fmt"

func main() {
	mensagemBoasVindas()
	peso := solicitaPeso()
	altura := solicitaAltura()
	realizandoCalculo(peso, altura)
}

func mensagemBoasVindas() {
	fmt.Println("*********************************")
	fmt.Println("*                               *")
	fmt.Println("*  Bem Vindo a Calculadora IMC  *")
	fmt.Println("*                               *")
	fmt.Println("*********************************")
}

func solicitaPeso() float32 {
	var peso float32
	fmt.Print("Digite o seu peso (em kg): ")
	fmt.Scan(&peso)
	return peso
}

func solicitaAltura() float32 {
	var altura float32
	fmt.Print("Digite sua altura (em metros): ")
	fmt.Scan(&altura)
	return altura
}

func validaPeso(peso float32) bool {
	if peso <= 0 || peso > 1000 {
		return false
	}
	return true
}

func validaAltura(altura float32) bool {
	if altura <= 0 || altura > 4 {
		return false
	}
	return true
}

func calculaIMC(peso float32, altura float32) float32 {
	return peso / (altura * altura)
}

func mensagemIMC(imc float32) {
	fmt.Printf("Seu IMC é: %.2f", imc)
}

func mensagemDeErro() {
	fmt.Println("Não foi possível calcular o seu IMC.")
}

func realizandoCalculo(peso float32, altura float32) {
	var pesoValido = validaPeso(peso)
	var alturaValida = validaAltura(altura)

	if pesoValido && alturaValida {
		var imc = calculaIMC(peso, altura)
		mensagemIMC(imc)
	} else {
		mensagemDeErro()
	}
}
