package main

import (
	"fmt"
	"math"
)

type figura interface {
	area() float64
}

type quadrado struct {
	lado float64
}
type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func info(f figura) float64 {
	return f.area()
}
func main() {
	quad := quadrado{4}
	circ := circulo{5}

	fmt.Println(info(quad))
	fmt.Println(info(circ))

}
