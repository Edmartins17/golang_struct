package model

type Vehicle struct {
	Year  int
	Plate string
	Model string
}

type Car struct {
	Vehicle
	Power              int
	Ports              int
	HasAirConditioning bool
}

type MotorCycle struct {
	Vehicle
	CylinderCapacity int
}
