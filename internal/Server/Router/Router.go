package Router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"main/internal/API/Controllers"
	"net/http"
)

func Router(router *mux.Router) {

	router.HandleFunc("/golang/api/v1/get_img", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		files := Controllers.GetWidgets()

		json.NewEncoder(w).Encode(files)

	}).Methods("GET", "OPTIONS")
	router.HandleFunc("/golang/api/v1/Upd_img", func(w http.ResponseWriter, r *http.Request) {
		status := Controllers.UpdWidgets()
		fmt.Fprintf(w, status)

	}).Methods("GET", "OPTIONS")
}
