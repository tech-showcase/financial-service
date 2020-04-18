package main

import (
	"fmt"
	"github.com/tech-showcase/financial-service/api"
	"github.com/tech-showcase/financial-service/cmd"
)

func main() {
	fmt.Println("Hi, I am Financial Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
