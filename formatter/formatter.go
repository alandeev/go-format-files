package formatter

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/alandev2/prettier/utils"
)

var wg sync.WaitGroup

type ExtensionOptions struct {
	match  string
	folder string
}

type Formatter struct {
	extensions []ExtensionOptions
}

func NewFormatter() *Formatter {
	formatter := &Formatter{
		extensions: []ExtensionOptions{
			{
				match:  "*.json",
				folder: "json",
			},
			{
				match:  "*.zip",
				folder: "zip",
			},
			{
				match:  "*.tar*",
				folder: "zip",
			},
			{
				match:  "*.png",
				folder: "png",
			},
			{
				match:  "*.jpg",
				folder: "jpg",
			},
			{
				match:  "*.csv",
				folder: "csv",
			},
			{
				match:  "*.deb",
				folder: "deb",
			},
			{
				match:  "*.txt",
				folder: "txt",
			},
			{
				match:  "*.pdf",
				folder: "pdf",
			},
			{
				match:  "*.pptx",
				folder: "pptx",
			},
		},
	}

	return formatter
}

func FormatFiles(match string, folder string) {
	files, err := filepath.Glob(match)
	if err != nil {
		panic(err)
	}

	if len(files) > 0 {
		if !utils.FileExists(folder) {
			os.Mkdir(folder, os.ModePerm)
		}

		for _, file := range files {
			info, _ := os.Stat(file)
			if !info.IsDir() {
				os.Rename(
					filepath.Join(".", file),
					filepath.Join(".", folder, file),
				)
			}
		}
	}

	wg.Done()
}

func (f *Formatter) Run() {
	wg.Add(len(f.extensions))
	for _, formatOption := range f.extensions {
		go FormatFiles(formatOption.match, formatOption.folder)
	}

	wg.Wait()
	go FormatFiles("*", "outros")

	fmt.Println("Arquivos formatados com sucesso!")
}
