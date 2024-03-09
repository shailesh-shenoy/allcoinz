package api

import (
	"encoding/json"
	"net/http"

	"github.com/shailesh-shenoy/allcoinz/domain"
)

// * User routes and handler functions

func (apiServer *ApiServer) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		HandleError(w, r, err)
		return
	}
	if err := apiServer.UserService.CreateUser(r.Context(), &user); err != nil {
		HandleError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	return

}
