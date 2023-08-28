package main

import "fmt"

type InputData struct {
	fuelin float32
	carType string
}

func (car InputData) jarak() float32{
	const fuel float32 = 1.5 
	calculate := car.fuelin/fuel
	
	return calculate
}

func main() {
	var car InputData
	car.carType="SUV"
	car.fuelin=10.5

	fmt.Println("car type : ", car.carType)
	fmt.Println("est. distance : ", car.jarak())
}