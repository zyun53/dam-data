package handler

import (
	"net/http"

	"github.com/zyun-i/dam-data/model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllDams(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	dams := []model.Dam{}
	db.Find(&dams)
	respondJSON(w, http.StatusOK, dams)
}

func GetDam(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	dam := getDamOr404(db, id, w, r)
	if dam == nil {
		return
	}

	respondJSON(w, http.StatusOK, dam)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getDamOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Dam {
	dam := model.Dam{}
	if err := db.First(&dam, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &dam
}
