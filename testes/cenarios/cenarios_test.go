package cenarios_test

import "testing"

func TestFuncao(t *testing.T) {
	t.Run("Testando um funcao qualquer", func(t *testing.T) {

	})
}

func TestCenarios(t *testing.T) {
	t.Run("Testando uma outra funcao qualquer", func(t *testing.T) {
		t.Run("Esse cenario deve pular", func(t *testing.T) {

		})
		t.Run("Esse cenario deve pular tambem", func(t *testing.T) {

		})
	})
}
