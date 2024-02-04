package vehicles

import (
	"errors"
)
type Vehicle interface {
	Distance() float64
}
const (
	carVehicle = "CAR"
	motorcycleVehicle = "MOTORCYCLE"
	truckVehicle = "TRUCK"
	planeVehicle = "PLANE"
)

func New( vehicle string, time int) (Vehicle, error) {

	switch vehicle {
	case carVehicle:
		return &Car{Time: time}, nil
	case motorcycleVehicle:
		return &Motocycle{Time: time}, nil
	case truckVehicle:
		return &Truck{Time: time}, nil
	case planeVehicle:
		return &Plane{Time: time}, nil
	default:
		return nil, errors.New("invalid vehicle type")
	}

}


type Car struct {
	Time int
}
func (c *Car) Distance() float64 {
	return 100* (float64(c.Time)/60)
}

type Motocycle struct {
	Time int
}

func (m *Motocycle) Distance() float64 {
	return 120* (float64(m.Time)/60)
}

type Truck struct {
	Time int
}

func (t *Truck) Distance() float64 {
	return 90* (float64(t.Time)/60)
}

type Plane struct {
	Time int
}

func (p *Plane) Distance() float64 {
	return 800* (float64(p.Time)/60)
}