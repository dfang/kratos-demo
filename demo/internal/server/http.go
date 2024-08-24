package server

import (
	v1 "demo/api/helloworld/v1"
	"demo/internal/conf"
	"demo/internal/service"

	// "demo/internal/server/assets"
	"demo"
	// rest "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerUI "github.com/tx7do/kratos-swagger-ui"

	nhttp "net/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	// swaggerUI.RegisterSwaggerUIServer(
	// 	srv,
	// 	"DEMO",
	// 	// "https://petstore3.swagger.io/api/v3/openapi.yaml",
	// 	http.FileServer(http.FS(assets.OpenApiData)),
	// 	"/docs/",
	// )

	// swaggerHandler := swaggerUI.NewWithOption(
	// 	swaggerUI.WithTitle("Petstore"),
	// 	swaggerUI.WithLocalFile(".\\openapi.yaml"),
	// 	swaggerUI.WithBasePath("/docs/"),
	// )
	// srv.Handle("/docs/", swaggerHandler)

	nhttp.Handle("/static/", nhttp.StripPrefix("/static/", nhttp.FileServer(nhttp.FS(demo.OpenApiFile))))
	swaggerHandler := swaggerUI.New(
		"Petstore",
		// "https://petstore3.swagger.io/api/v3/openapi.json",
		"/static/openapi.yaml",
		"/docs/",
	)
	srv.HandlePrefix("/docs/", swaggerHandler)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
