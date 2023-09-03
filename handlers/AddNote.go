package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"notes-api/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
	}

	var note models.Note
	err = json.Unmarshal(body, &note)

	if result := h.DB.Create(&note); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send response indicating note was created
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Note created")

}
