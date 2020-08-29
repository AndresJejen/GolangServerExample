package test

import (
	"testing"

	"github.com/AndresJejen/AprendiendoGo/calculadora"
	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	resultado := calculadora.Sumar(12.5, 6.8)
	assert.Equal(t, float32(19.3), resultado)
}

func TestRestar(t *testing.T) {
	resultado := calculadora.Restar(12.5, 5)
	assert.Equal(t, float32(7.5), resultado)
}

func TestMultiplicar(t *testing.T) {
	resultado := calculadora.Multiplicar(12.5, 5)
	assert.Equal(t, float32(62.5), resultado)
}

func TestDividirOK(t *testing.T) {
	resultado, err := calculadora.Dividir(12.5, 5)
	assert.Equal(t, float32(2.5), resultado)
	assert.Nil(t, err)
}

func TestDividirFail(t *testing.T) {
	resultado, err := calculadora.Dividir(12.5, 0)
	assert.Equal(t, float32(0), resultado)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "NO es posible dividir por 0")
}
