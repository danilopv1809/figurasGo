package figuras

import "fmt"

type Cuadrado struct {
	Ancho  float64
	Altura float64
}

func (*Cuadrado) getName() string {
	return "Cuadrado"
}
func (c *Cuadrado) cArea() string {
	area := c.Ancho * c.Altura
	return fmt.Sprintf("Figura -> %s\n, Area: %.2f", c.getName(), area)
}
func (c *Cuadrado) cPerimetro() string {
	perimetro := (2 * c.Ancho) + (2 * c.Altura)
	return fmt.Sprintf("Figura -> %s\n, Area: %.2f", c.getName(), perimetro)
}
