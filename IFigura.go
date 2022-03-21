package figuras

type IFigura interface {
	getName() string
	cArea() string
	cPerimetro() string
}

func ToString(figura IFigura) string {
	name := figura.getName()
	area := figura.cArea()
	perimetro := figura.cPerimetro()
	return name + "\n" + area + "\n" + perimetro
}
