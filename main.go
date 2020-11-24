package main

import (
	"fmt"
	"github.com/tech-showcase/financial-service/api"
	"github.com/tech-showcase/financial-service/cmd"
	"github.com/tech-showcase/financial-service/config"
)

func init() {
	config.Instance = config.Read()
}

func main() {
	fmt.Println("Hi, I am Financial Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
