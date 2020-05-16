package global

import (
	"github.com/tech-showcase/financial-service/config"
)

var Configuration = config.Config{}

func init() {
	Configuration = config.Read()
}
