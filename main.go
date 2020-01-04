package goarea

import "math"

// Pi é uma proporção nuémira definida pela relação entre
// o perimetro de uma circunferencia e seu diâmetro
const Pi = 3.1416

// Circ é responsável por calcular a area da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por caçcular a area da circuferência
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
