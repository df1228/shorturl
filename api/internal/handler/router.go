package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/pathvar"
)

type ginRouter struct {
	g *gin.Engine
}

// NewRouter returns a gin.Router.
func NewRouter(opts ...Option) httpx.Router {
	g := gin.New()
	cfg := config{
		redirectTrailingSlash: true,
		redirectFixedPath:     false,
	}
	cfg.options(opts...)

	g.RedirectTrailingSlash = cfg.redirectTrailingSlash
	g.RedirectFixedPath = cfg.redirectFixedPath
	return &ginRouter{g: g}
}

func (pr *ginRouter) Handle(method, reqPath string, handler http.Handler) error {
	if !validMethod(method) {
		return ErrInvalidMethod
	}

	if len(reqPath) == 0 || reqPath[0] != '/' {
		return ErrInvalidPath
	}

	pr.g.Handle(strings.ToUpper(method), reqPath, func(ctx *gin.Context) {
		params := make(map[string]string)
		for i := 0; i < len(ctx.Params); i++ {
			params[ctx.Params[i].Key] = ctx.Params[i].Value
		}
		if len(params) > 0 {
			ctx.Request = pathvar.WithVars(ctx.Request, params)
		}
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	})
	return nil
}

func (pr *ginRouter) SetNotFoundHandler(handler http.Handler) {
	pr.g.NoRoute(gin.WrapH(handler))
}

func (pr *ginRouter) SetNotAllowedHandler(handler http.Handler) {
	pr.g.NoMethod(gin.WrapH(handler))
}

func (pr *ginRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pr.g.ServeHTTP(w, r)
}

func validMethod(method string) bool {
	return method == http.MethodDelete || method == http.MethodGet ||
		method == http.MethodHead || method == http.MethodOptions ||
		method == http.MethodPatch || method == http.MethodPost ||
		method == http.MethodPut
}

var (
	// ErrInvalidMethod is an error that indicates not a valid http method.
	ErrInvalidMethod = errors.New("not a valid http method")
	// ErrInvalidPath is an error that indicates path is not start with /.
	ErrInvalidPath = errors.New("path must begin with '/'")
)

type config struct {
	redirectTrailingSlash bool
	redirectFixedPath     bool
}

func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
}

type Option func(o *config)

func WithRedirectTrailingSlash(redirectTrailingSlash bool) Option {
	return func(c *config) {
		c.redirectTrailingSlash = redirectTrailingSlash
	}
}

func WithRedirectFixedPath(redirectFixedPath bool) Option {
	return func(c *config) {
		c.redirectFixedPath = redirectFixedPath
	}
}
