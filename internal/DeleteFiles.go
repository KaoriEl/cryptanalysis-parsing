package internal

import (
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func DeleteAll() {
	filePrefix, _ := filepath.Abs("/var/www/investments-cryptanalysis-parsing/assets/img") // path from the working directory
	dir, err := ioutil.ReadDir(filePrefix)
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to read dir"))
	}
	for _, d := range dir {
		err := os.RemoveAll(path.Join([]string{filePrefix, d.Name()}...))
		if err != nil {
			color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to delete files"))
		}
	}
}
