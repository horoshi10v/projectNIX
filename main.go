package main
import "fmt"
import "github.com/google/uuid"

func main() {
	fmt.Printf("Hello NIX!")
	id := uuid.New()
	fmt.Printf(id.String())
=======
import (
	"fmt"
	"math"
)

func distance(x, y, x2, y2 float64) float64 {
	a := x2 - x
	b := y2 - y
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) Area() float64 {
	l := distance(r.x, r.y, r.x, r.y2)
	w := distance(r.x, r.y, r.x2, r.y)
	return l * w
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type Position struct {
	x, y float64
}

type Rectangle struct {
	Position //x and y
	x2, y2   float64
}
type Circle struct {
	Position
	r float64
}

type Shape interface {
	Area() float64
}

func totalArea(shape ...Shape) (total float64) {
	for _, shape := range shape {
		total += shape.Area()
	}
	return total
}

func main() {

	rectangle := &Rectangle{
		Position: Position{
			x: 0,
			y: 0,
		},
		x2: 10,
		y2: 10,
	}

	circle := &Circle{
		r: 5,
	}
	circle2 := &Circle{
		r: 5,
	}

	fmt.Println(rectangle.Area())
	fmt.Println(circle.Area())
	fmt.Println(totalArea(circle, rectangle, circle2, circle2))

}
