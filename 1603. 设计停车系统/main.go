package main

import "fmt"

type ParkingSystem struct {
	car map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{car: map[int]int{1: big, 2: medium, 3: small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	this.car[carType] = this.car[carType] - 1

	if this.car[carType] < 0 {
		return false
	}
	return true
}

func main() {
	big := 1
	medium := 1
	small := 0

	obj := Constructor(big, medium, small)
	fmt.Println(obj.AddCar(1))
	fmt.Println(obj.AddCar(2))
	fmt.Println(obj.AddCar(3))
	fmt.Println(obj.AddCar(1))
}
