package global

import (
	"github.com/tech-showcase/financial-service/config"
)

var Configuration = config.Config{}

func init() {
	var err error
	Configuration, err = config.Read()
	if err != nil {
		panic(err)
	}
}
