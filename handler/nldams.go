package handler

import (
	"net/http"

	"github.com/zyun-i/dam-data/model"

	"gorm.io/gorm"
)

func GetAllNlDams(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	dams := []model.NlDam{}
	//db.Find(&dams)
	//db.Table("nl_dams").Select("nl_dams.*, nl_dam_types.*").Joins("left join nl_dam_types on nl_dam_types.id = nl_dams.type").Scan(&dams)
	db.Joins("NlDamType").Find(&dams)
	respondJSON(w, http.StatusOK, dams)
}
