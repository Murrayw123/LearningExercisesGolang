package OOP

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base * 3
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (c Circle) Perimeter() float64 {
	return c.Radius * 2 * 3.14
}

func OOPExerciseOne() {
	triangleArea := Triangle{Base: 3, Height: 4}.Area()
	trianglePerimeter := Triangle{Base: 3, Height: 4}.Perimeter()

	circleArea := Circle{Radius: 5}.Area()
	circlePerimeter := Circle{Radius: 5}.Perimeter()

	rectangleArea := Rectangle{Width: 3, Height: 4}.Area()
	rectanglePerimeter := Rectangle{Width: 3, Height: 4}.Perimeter()

	fmt.Println("Triangle Area: ", triangleArea)
	fmt.Println("Triangle Perimeter: ", trianglePerimeter)

	fmt.Println("Circle Area: ", circleArea)
	fmt.Println("Circle Perimeter: ", circlePerimeter)

	fmt.Println("Rectangle Area: ", rectangleArea)
	fmt.Println("Rectangle Perimeter: ", rectanglePerimeter)
}
