package main

import "fmt"

// Giả sử ta có một lớp hình học có phương thức vẽ có thẻ vẽ theo chiều ngang hoặc chiều đứng
// Tối ưu là tạo ra 1 lớp DrawContext thay vì tạo ra 2 lớp là ShapeWithDrawHorizotal và ShapeWithDrawVertical

type ShapeInterface interface {
	Draw()
}

type DrawerInterface interface {
	Draw()
}

type DrawHorizotal struct{}

func (s *DrawHorizotal) Draw() {
	fmt.Println("Draw Horizotal")
}

type DrawVertical struct{}

func (s *DrawVertical) Draw() {
	fmt.Println("Draw Vertical")
}

type Shape struct {
	DrawContext DrawerInterface
}

func (s *Shape) Draw() {
	s.DrawContext.Draw()
}

type Circle struct {
	Shape
	Radius float64
}

func (s *Circle) Draw() {
	fmt.Printf("Circle with Radius %v ", s.Radius)
	s.Shape.Draw()
}

type Square struct {
	Shape
	Width float64
}

func (s *Square) Draw() {
	fmt.Printf("Square with Width %v ", s.Width)
	s.Shape.Draw()
}

func main() {
	drawerHoz := DrawHorizotal{}
	drawerVer := DrawVertical{}

	cCircle := Circle{Radius: 0.5}
	cCircle.DrawContext = &drawerHoz
	cCircle.Draw()

	cCircle.DrawContext = &drawerVer
	cCircle.Draw()

	aSquare := Square{Width: 10}
	aSquare.DrawContext = &drawerHoz
	aSquare.Draw()

	aSquare.DrawContext = &drawerVer
	aSquare.Draw()

}
