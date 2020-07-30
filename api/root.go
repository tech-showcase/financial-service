package api

import (
	"fmt"
	"github.com/soheilhy/cmux"
	"net"
)

func Activate(port int) {
	lisAddress := fmt.Sprintf(":%d", port)
	lis, err := net.Listen("tcp", lisAddress)
	if err != nil {
		panic(err)
	}

	m := cmux.New(lis)
	RegisterFinancialGRPCAPI(m)
	RegisterFinancialHTTPAPI(m)
	m.Serve()
}
