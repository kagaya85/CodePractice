package main

type ParkingSystem struct {
	conut [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{conut: [4]int{0, big, medium, small}}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if p.conut[carType] > 0 {
		p.conut[carType]--
		return true
	}
	return false
}
