package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteNote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// new var reference to note struct to be deleted
	var note models.Note

	// Try to delete note
	if result := h.DB.Delete(&note, id); result.Error != nil {
		fmt.Println("Error while trying to delete: ", result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send response indicating note was deleted
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Note deleted")
}
