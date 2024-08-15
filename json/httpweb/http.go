package httpweb

import (
	"encoding/json"
	"json-golang/domain"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Age += 1
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
