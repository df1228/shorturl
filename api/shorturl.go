package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dfang/shorturl/api/internal/config"
	"github.com/dfang/shorturl/api/internal/handler"
	"github.com/dfang/shorturl/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/shorturl-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	r := handler.NewRouter()
	engine := rest.MustNewServer(
		c.RestConf,
		// rest.RestConf{
		// 	ServiceConf: service.ServiceConf{
		// 		Log: logx.LogConf{
		// 			Mode: "console",
		// 		},
		// 	},
		// 	Port:     c.Port,
		// 	Timeout:  20000,
		// 	MaxConns: 500,
		// },
		rest.WithRouter(r))

	// NotFound defines a handler to respond whenever a route could not be found.
	r.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("nothing here"))
	}))

	// MethodNotAllowed defines a handler to respond whenever a method is not allowed.
	r.SetNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("405"))
	}))
	defer engine.Stop()

	// server := rest.MustNewServer(c.RestConf, rest.WithCors())
	// defer server.Stop()
	// handler.RegisterHandlers(server, ctx)
	handler.RegisterHandlers(engine, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	engine.Start()

	// fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// server.Start()
}
