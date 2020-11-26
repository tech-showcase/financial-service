package main

import (
	"fmt"
	"github.com/tech-showcase/financial-service/api"
	"github.com/tech-showcase/financial-service/cmd"
	"github.com/tech-showcase/financial-service/config"
	"github.com/tech-showcase/financial-service/helper"
)

func init() {
	var err error
	config.Instance, err = config.Read()
	if err != nil {
		panic(err)
	}

	helper.OAuth2HelperInstance = helper.NewOAuth2Helper()
}

func main() {
	fmt.Println("Hi, I am Financial Service!")

	args := cmd.Parse()

	api.Activate(args.Port)
}
