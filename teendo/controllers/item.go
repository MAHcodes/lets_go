package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MAHcodes/lets_go/teendo/database"
	"github.com/MAHcodes/lets_go/teendo/models"
	"github.com/MAHcodes/lets_go/teendo/utils"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := database.DB.First(&item, r.PathValue("id")).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Item not found")
		return
	}
	utils.RespondJSON(w, http.StatusOK, item)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	database.DB.Find(&items)
	utils.RespondJSON(w, http.StatusOK, items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	database.DB.Create(&item)
	utils.RespondJSON(w, http.StatusCreated, item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := database.DB.First(&item, r.PathValue("id")).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Item not found")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	database.DB.Save(&item)
	utils.RespondJSON(w, http.StatusOK, item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := database.DB.First(&item, r.PathValue("id")).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, "Item not found")
		return
	}
	database.DB.Delete(&item)
	utils.RespondJSON(w, http.StatusNoContent, nil)
}
