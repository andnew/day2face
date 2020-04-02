package ch108

// 结构体中嵌入接口

import (
	"fmt"
)

func MainCall3() {
	demo31()
}

type HornSounder interface {
	SoundHorn()
}

type Vehicle struct {
	List [2]HornSounder
}

type Car struct {
	Sound string
}

type Bike struct {
	Sound string
}

func demo31() {
	vehicle := new(Vehicle)
	vehicle.List[0] = &Car{"BEEP"}
	vehicle.List[1] = &Bike{"RING"}

	for _, hornSounder := range vehicle.List {
		hornSounder.SoundHorn()
		// PressHorn(hornSounder) 这种方式也可以
	}
}

func (car *Car) SoundHorn() {
	fmt.Println(car.Sound)
}

func (bike *Bike) SoundHorn() {
	fmt.Println(bike.Sound)
}

func PressHorn(hornSounder HornSounder) {
	hornSounder.SoundHorn()
}
