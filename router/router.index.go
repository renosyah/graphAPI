package router

import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}
