package tools

import (
	"net/http"
)

func Headers(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "text/html; charset=UTF-8")
	(*w).Header().Set("Connection", "Keep-Alive")
	(*w).Header().Set("Keep-Alive", "300")
	// finish header
	(*w).WriteHeader(http.StatusOK)
}
