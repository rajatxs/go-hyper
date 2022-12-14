package server

import "net/http"

/*
Writes CORS headers to ResponseWriter

Accept next http handler function
*/
func useCORS(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, OPTIONS")

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, X-Public-Id, X-Public-Key")
		next.ServeHTTP(w, r)
	})
}
