package shapes

import (
	"errors"
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

var ErrNegative = errors.New("negative radius")

func NewCircle(r float64) (*Circle, error) {
	if r < 0 {
		return nil, fmt.Errorf("new circle: %w", ErrNegative)
	}

	return &Circle{Radius: r}, nil
}

func NewRectangle(w, h float64) (*Rectangle, error) {
	if w < 0 || h < 0 {
		return nil, fmt.Errorf("new rectangle: %w", ErrNegative)
	}

	return &Rectangle{Width: w, Height: h}, nil
}


func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func (c *Circle) String() string {
	return fmt.Sprintf("circle(r=%.2f)", c.Radius)
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}