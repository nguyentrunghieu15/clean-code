package main

import "fmt"

type TransportInterface interface {
	Move()
}

type TransportByCar struct {
	Velocity float64
}

func (c *TransportByCar) Move() {
	fmt.Println("Move by car with velocity: ", c.Velocity)
}

type TransportByTrain struct {
	Velocity float64
}

func (c *TransportByTrain) Move() {
	fmt.Println("Move by train with velocity: ", c.Velocity)
}

type TransportMethod int

const (
	BYCAR = iota
	BYTRAIN
)

func getTransport(method TransportMethod, velocity float64) TransportInterface {
	switch method {
	case BYCAR:
		return &TransportByCar{Velocity: velocity}
	case BYTRAIN:
		return &TransportByTrain{Velocity: velocity}
	}
	return nil
}

func main() {
	usingCar := getTransport(BYCAR, 60)
	usingCar.Move()
	usingTrain := getTransport(BYTRAIN, 60)
	usingTrain.Move()
}
