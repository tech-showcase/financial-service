package main

import (
	"github.com/tech-showcase/financial-service/api"
	"github.com/tech-showcase/financial-service/cmd"
)

func main() {
	args := cmd.Parse()

	api.Activate(args.Port)
}
