package main

import (
	"fmt"
	"github.com/tech-showcase/financial-service/api"
	"github.com/tech-showcase/financial-service/cmd"
	"github.com/tech-showcase/financial-service/config"
)

func init() {
	var err error
	config.Instance, err = config.Read()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hi, I am Financial Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
