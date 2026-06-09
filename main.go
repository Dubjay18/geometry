package main

import (
	"errors"
	"fmt"

	"github.com/Dubjay18/geometry/shapes"
)

func main() {
	circle, err := shapes.NewCircle(5)
	if err != nil {
		panic(err)
	}
	rectangle, err := shapes.NewRectangle(10, 5)
	if err != nil {
		panic(err)
	}
	shps := []shapes.Shape{
		circle,
		rectangle,
	}

	for _, s := range shps {
		fmt.Println("Area:", s.Area())
		fmt.Println("Perimeter:", s.Perimeter())
		switch v := s.(type) {
		case shapes.Circle:
			fmt.Println("This is a circle with radius", v.Radius)
		case shapes.Rectangle:
			fmt.Println("This is a rectangle with width", v.Width, "and height", v.Height)
		}
	}

	_, err = shapes.NewCircle(-1)
	if errors.Is(err, shapes.ErrNegative) {
		fmt.Println("Cannot create circle with negative radius")
	}

}
