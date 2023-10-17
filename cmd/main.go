package main

import "github.com/MikhailGulkin/SimpleBlockChainSystemExample/internal/api"

func main() {
	err := api.Run()
	if err != nil {
		panic(err)
	}
}
