package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

type quadrado struct {
	lado float64
}

func (r retangulo) area() float64 { //r referencia struct que possui largura e altura (array de argumentos)
	return r.largura * r.altura
}

func (r retangulo) perimetro() float64 {
	return (r.largura * 2) + (r.altura * 2)
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

func (q quadrado) perimetro() float64 {
	return q.lado * 4
}

func measure(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	ret := retangulo{
		3,
		4,
	}
	circ := circulo{
		3,
	}
	quad := quadrado{
		5,
	}
	measure(ret)
	measure(circ)
	measure(quad)
}
