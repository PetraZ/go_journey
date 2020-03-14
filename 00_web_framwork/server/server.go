package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

//Params are the params to initialize the server
type Params struct {
	fx.In

	R  *gin.Engine
	Lc fx.Lifecycle
}

//NewServer is creator to initialize the server
func NewServer(p Params) *http.Server {
	host := "localhost"
	port := 8100
	addr := fmt.Sprintf("%s:%d", host, port)

	server := &http.Server{
		Addr:    addr,
		Handler: p.R,
	}
	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					fmt.Println("something wrong")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return server
}
