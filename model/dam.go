package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Dam struct {
	gorm.Model

	Name   string
	City   string
	State  string
	Street string

	Latitude  string
	Longitude string

	WaterSystem string
	RiverName   string

	Type    string
	Purpose string

	BankHeight    int
	BankLength    int
	BankVolume    int
	TotalCapacity int

	InstitutionInCharge int
	YearOfCompletion    int

	OsmId       int
	OsmElements string
}
