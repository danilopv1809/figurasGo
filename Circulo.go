package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float64
}

func (*Circulo) getName() string {
	return "Circulo"
}
func (c *Circulo) cArea() string {
	area := math.Pi * (c.Radio * c.Radio)
	return fmt.Sprintf("Figura -> %s\n, Area: %.2f", c.getName(), area)
}
func (c *Circulo) cPerimetro() string {
	perimetro := 2 * math.Pi * c.Radio
	return fmt.Sprintf("Figura -> %s\n, Area: %.2f", c.getName(), perimetro)
}
