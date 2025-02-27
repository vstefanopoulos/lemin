package main

import (
	"lemin/lemin"
	"log"
)

func main() {
	input, result, err := lemin.Please()
	if err != nil {
		log.Fatal(err)
	}

	lemin.PrintResult(input, result)
}
