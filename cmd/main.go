package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("Wrong use here")
	}

	path := args[1]
	read(path)
}

func read(path string) {
	out, err := exec.Command("ocrad", path).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
