package cmd

import (
	"flag"
	"github.com/tech-showcase/financial-service/presenter"
)

func Parse() (args presenter.Args) {
	flag.IntVar(&args.Port, "port", 8080, "Port which apps will listen to")
	flag.Parse()

	return
}
