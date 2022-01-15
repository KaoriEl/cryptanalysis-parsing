package Router

import (
	"fmt"
	"github.com/gorilla/mux"
	"main/internal/API/Controllers"
	"net/http"
)

func Router(router *mux.Router) {
	router.HandleFunc("/api/v1/get_img", func(w http.ResponseWriter, r *http.Request) {
		files := Controllers.GetWidgets()
		fmt.Fprintf(w, string(files))

	})
	router.HandleFunc("/api/v1/Upd_img", func(w http.ResponseWriter, r *http.Request) {
		status := Controllers.UpdWidgets()
		fmt.Fprintf(w, status)

	})
}
