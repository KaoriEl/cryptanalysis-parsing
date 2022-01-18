package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

type Files struct {
	Files []string
}

// ClearDir Очищает заданную директорию от всех файлов кроме папки tmp/**
func ClearDir(pathFiles string) {
	filePrefix, _ := filepath.Abs(pathFiles) // path from the working directory
	dir, err := ioutil.ReadDir(filePrefix)
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to read dir"))
	}
	for _, d := range dir {
		if d.Name() != "tmp" && d.Name() != ".gitkeep" {
			err := os.RemoveAll(path.Join([]string{filePrefix, d.Name()}...))
			if err != nil {
				color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Failed to delete files"))
			}
		}

	}
}

// OnlyReadDir Только читает файлы и выводит их список/**
func OnlyReadDir(path string) Files {
	filePrefix, _ := filepath.Abs(path) // path from the working directory
	files, err := ioutil.ReadDir(filePrefix)
	if err != nil {
		log.Fatal(err)
	}
	var ctx Files
	for _, f := range files {
		if f.Name() != "tmp" && f.Name() != ".gitkeep" {
			ctx.Files = append(ctx.Files, "/golang/assets/img/"+f.Name())
		}
	}
	return ctx
}

// ReadDirAndCopy Копирует файлы из одной директории в другую/**
func ReadDirAndCopy(TakePath string, WritePath string) {
	TakefilePrefix, _ := filepath.Abs(TakePath)   // path from the working directory
	WritefilePrefix, _ := filepath.Abs(WritePath) // path from the working directory
	files, err := ioutil.ReadDir(TakefilePrefix)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() != ".gitkeep" {
			input, err := ioutil.ReadFile(TakefilePrefix + "/" + f.Name())
			if err != nil {
				fmt.Println(err)
				return
			}
			err = ioutil.WriteFile(WritefilePrefix+"/"+f.Name(), input, 0644)
			if err != nil {
				fmt.Println("Error creating", WritefilePrefix)
				fmt.Println(err)
				return
			}
		}

	}
}
