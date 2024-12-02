package main

import(
	"fmt"
	"math"
)

type Point struct{
	x float64
	y float64
}

func NewPoint(x,y float64) *Point{
	return &Point{x:x, y:y}
}

func (p *Point) GetX() float64{
	return p.x
}

func (p *Point) GetY() float64{
	return p.y
}

func lenght(p1,p2 *Point) float64{
	return math.Sqrt(math.Pow(p2.GetX() - p1.GetX(),2)+ math.Pow(p2.GetY()-p1.GetY(),2))
}

func main(){
	point1 := NewPoint(1.0, 2.0)
	point2 := NewPoint(4.0, 6.0)

	fmt.Println(lenght(point1, point2))
}

