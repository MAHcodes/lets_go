package server

import (
	"fmt"
	"net/http"

	"github.com/MAHcodes/lets_go/teendo/config"
	"github.com/MAHcodes/lets_go/teendo/routes"
)

func NewServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%v", config.PORT),
		Handler: routes.RegisterRoutes(),
	}
}
