package formas

import "testing"

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f",
				areaRecebida,
				areaEsperada,
			)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(314.1592653589793)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f",
				areaRecebida,
				areaEsperada,
			)
		}
	})

}
