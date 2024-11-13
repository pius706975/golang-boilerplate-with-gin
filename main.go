package main

import (
	"go-gin/cmd"
	"log"
	"os"
)

func main() {
	err := cmd.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}