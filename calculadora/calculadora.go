package calculadora

import "errors"

func Sumar(numero1 float32, numero2 float32) float32 {
	return numero1 + numero2
}

func Restar(numero1 float32, numero2 float32) float32 {
	return numero1 - numero2
}

func Multiplicar(numero1 float32, numero2 float32) float32 {
	return numero1 * numero2
}

func Dividir(numero1 float32, numero2 float32) (float32, error) {
	if numero2 != 0 {
		return numero1 / numero2, nil
	}
	return 0, errors.New("NO es posible dividir por 0")
}
