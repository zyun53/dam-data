package model

import (
	"gorm.io/gorm"
)

type NlDam struct {
	gorm.Model
	Code                           int32            `json:"code"`
	PositionLat                    float32          `json:"position_lat"`
	PositionLng                    float32          `json:"position_lng"`
	Name                           string           `json:"name"`
	WaterSystemName                string           `json:"water_system_name"`
	RiverName                      string           `json:"river_name"`
	Type                           NlDamType        `gorm:"embedded"`
	Purpose                        []NlDamPurpose   `gorm:"many2many:nl_dam_nl_purposes;"`
	DamScaleBankHeight             int32            `json:"dam_scale_bank_height"`
	DamScaleBankSpan               int32            `json:"dam_scale_bank_span"`
	BankVolume                     int32            `json:"bank_volume"`
	TotalPondage                   int32            `json:"total_pondage"`
	InstitutionInCharge            NlDamInstitution `gorm:"embedded"`
	YearOfCompletion               int32            `json:"year_of_completion"`
	Address                        string           `json:"address"`
	PositionalInformationPrecision NlDamPrecision   `gorm:"embedded"`
}

// ダム型式コード
type NlDamType struct {
	ID    int    `json:"dam_type_id" gorm:"primaryKey"`
	Value string `json:"dam_type_name"`
}

// ダム目的コード
type NlDamPurpose struct {
	ID    int `gorm:"primaryKey"`
	Value string
}

// 位置情報精度コード
type NlDamPrecision struct {
	ID    int `gorm:"primaryKey"`
	Value string
}

// ダム事業者コード
type NlDamInstitution struct {
	ID    int `gorm:"primaryKey"`
	Value string
}
