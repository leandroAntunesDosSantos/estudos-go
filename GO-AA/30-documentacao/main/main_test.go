package main

import "testing"

// soma(i ...int) int
func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expexted:", resultado, "Got:", teste)
	}
}

func TestSoma2(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Error("Expexted:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != 100 {
		t.Error("Expexted:", resultado, "Got:", teste)

	}
}
