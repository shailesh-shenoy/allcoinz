package api

import (
	"encoding/json"
	"net/http"

	"github.com/shailesh-shenoy/allcoinz/domain"
)

type ApiServer struct {
	ListenAddr string
	mux        *http.ServeMux

	// Service interfaces to be used by various routes
	UserService domain.UserService
}

func (s *ApiServer) Run() (err error) {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		// return hello world json response

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, mux world!"}`))

	})
	s.mux.HandleFunc("POST /users", s.handleUserCreate)

	return http.ListenAndServe(s.ListenAddr, s.mux)

}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(&err)
}
