package model

import (
	"gorm.io/gorm"
)

type Dammap struct {
	gorm.Model
	DamNo         string  `json:"dam_no"`
	Name          string  `json:"name"`
	Kana          string  `json:"kana"`
	SaganSyozaiti string  `json:"sagan_syozaiti"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	McLat         float64 `json:"mc_lat"`
	McLng         float64 `json:"mc_lng"`
	Suikei        string  `json:"suikei"`
	Kasen         string  `json:"kasen"`
	Kou           float64 `json:"kou"`
	Tyou          float64 `json:"tyou"`
	Tai           float64 `json:"tai"`
	Ryuuiki       float64 `json:"ryuuiki"`
	Tansui        float64 `json:"tansui"`
	Tyokusetsu    float64 `json:"tyokusetsu"`
	Kansetsu      float64 `json:"kansetsu"`
	Sou           float64 `json:"sou"`
	Yuukou        float64 `json:"yuukou"`
	Jigyousya     string  `json:"jigyousya"`
	Sekousya      string  `json:"sekousya"`
	Tyakusyu      int32   `json:"tyakusyu"`
	Syunkou       int32   `json:"syunkou"`
	DamkoMei      string  `json:"damkomei"`
	DamkoKana     string  `json:"damkokana"`
	AreaCode      string  `json:"area_code"`
	TenkiPref     string  `json:"tenki_pref"`
	LocationName  string  `json:"location_name"`
	LocationUrl   string  `json:"location_url"`
	DataId        int64   `json:"data_id"`
	PrefNo        int64   `json:"pref_no"`
	KeishikiCode  string  `json:"keishiki_code"`
	PrefName      string  `json:"pref_name"`
	KeishikiName  string  `json:"keishiki_name"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&Dam{},
		&Dammap{},
	)
	return db
}
