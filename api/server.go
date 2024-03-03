package api

import "net/http"

type ApiServer struct {
	ListenAddr string
	mux        *http.ServeMux
}

func (s *ApiServer) Run() {
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		// return hello world json response

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello, mux world!"}`))
	})
	http.ListenAndServe(s.ListenAddr, s.mux)
}
