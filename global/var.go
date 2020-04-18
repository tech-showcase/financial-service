package global

import (
	"github.com/tech-showcase/financial-service/config"
	"github.com/tech-showcase/financial-service/presenter"
)

var Configuration = presenter.Config{}

func init() {
	var err error
	Configuration, err = config.Parse()
	if err != nil {
		panic(err)
	}
}
