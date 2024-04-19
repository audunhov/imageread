package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Reader interface {
	Read() (string, error)
}

type PngReader struct {
	path string
}

func (pngReader PngReader) Read() (string, error) {

	out, err := exec.Command("ocrad", pngReader.path).Output()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", out), nil
}

func getReader(path string) (Reader, error) {
	extension := filepath.Ext(path)

	if extension != ".png" {
		return nil, errors.New("Ivalid file extension")
	}

	return PngReader{
		path: path,
	}, nil
}

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("Invalid number of arguments")
	}

	path := args[1]

	reader, err := getReader(path)

	if err != nil {
		log.Fatal(err)
	}

	output, err := reader.Read()

	fmt.Printf("%s", output)
}
