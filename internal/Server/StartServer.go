package Server

import (
	"flag"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"main/internal/Extensions"
	"main/internal/Server/Router"
	"net/http"
	"path/filepath"
	"time"
)

func StartServer() {
	color.New(color.FgHiWhite).Add(color.Underline).Println("Server Tuning... ")
	Extensions.Progress(50)
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()
	filePrefix, _ := filepath.Abs("/var/www/investments-cryptanalysis-parsing/assets/img")
	dir = filePrefix
	router.PathPrefix("/golang/assets/img/").Handler(http.StripPrefix("/golang/assets/img/", http.FileServer(http.Dir(dir))))

	Router.Router(router)

	color.New(color.FgHiWhite).Add(color.Underline).Println("Start server. Port:3001 ")

	srv := &http.Server{
		Handler: router,
		Addr:    ":3001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
