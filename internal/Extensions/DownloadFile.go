package Extensions

import (
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Index(ImageUrl string, fileName string) {
	filePrefix, _ := filepath.Abs("/var/www/investments-cryptanalysis-parsing/assets/img/tmp") // path from the working directory
	fileName = filePrefix + fileName
	URL := ImageUrl
	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	color.New(color.FgGreen).Add(color.Bold).Println("File download in " + fileName)

}

func downloadFile(URL, fileName string) error {

	//Get the response bytes from the url

	go Progress(9531)

	response, err := http.Get(URL)

	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to get file by GET request"))
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to close request body"))
		}
	}(response.Body)

	if response.StatusCode != 200 {
		color.New(color.FgRed).Add(color.Underline).Println(errors.New("Received non 200 response code"))
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)

	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to create intermediate file"))
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to close file"))
		}
	}(file)
	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)

	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to write bytes to file"))
		return err
	}
	return nil
}
