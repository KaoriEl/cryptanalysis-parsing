package internal

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"net/http"
)

func StartServer() {
	color.New(color.FgHiWhite).Add(color.Underline).Println("Start server. Port:3002 ")
	http.HandleFunc("/hello", getFiles)
	if err := http.ListenAndServe(":3002", nil); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to start server"))
	}

}

func getFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asdads")
}
