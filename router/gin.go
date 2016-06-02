package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	gin *gin.Engine
}

func NewGin() Engine {
	gin.SetMode(gin.ReleaseMode)
	return ginEngine{gin: gin.New()}
}

func (r ginEngine) Debug(d bool) {
	if d {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Engine implementation

func (r ginEngine) Use() {
}

func (r ginEngine) Router() Router {
	return r
}

func (r ginEngine) Run(addr string) error {
	return r.gin.Run(addr)
}

// Router implementation

func (r ginEngine) GET(path string, h func(Response, Request) error) {
	r.gin.GET(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) POST(path string, h func(Response, Request) error) {
	r.gin.POST(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) DELETE(path string, h func(Response, Request) error) {
	r.gin.DELETE(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) PATCH(path string, h func(Response, Request) error) {
	r.gin.PATCH(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) PUT(path string, h func(Response, Request) error) {
	r.gin.PUT(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) OPTIONS(path string, h func(Response, Request) error) {
	r.gin.OPTIONS(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) HEAD(path string, h func(Response, Request) error) {
	r.gin.HEAD(path, func(c *gin.Context) {
		h(GinResponse{c}, GinRequest{c})
	})
}

func (r ginEngine) Handle(pattern string, handler http.Handler) {
	panic("gin Handle not yet implemented") // TODO
}

func (r ginEngine) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	panic("gin HandleFunc not yet implemented") // TODO
}

// Request implementation
type GinRequest struct {
	c *gin.Context
}

func (r GinRequest) Param(key string) string {
	return r.c.Param(key)
}

func (r GinRequest) Query(key string) string {
	return r.c.Query(key)
}

// Response implementation

type GinResponse struct {
	c *gin.Context
}

func (r GinResponse) AbortWithStatus(code int) error {
	r.c.AbortWithStatus(code)
	return nil
}

func (r GinResponse) String(code int, s string, v ...interface{}) error {
	r.c.String(code, s, v...)
	return nil
}

func (r GinResponse) HTML(code int, template string, v interface{}) error {
	r.c.HTML(code, template, v)
	return nil
}

func (r GinResponse) JSON(code int, obj interface{}) error {
	r.c.JSON(code, obj)
	return nil
}

func (r GinResponse) XML(code int, obj interface{}) error {
	r.c.XML(code, obj)
	return nil
}

func (r GinResponse) YAML(code int, obj interface{}) error {
	r.c.YAML(code, obj)
	return nil
}

func (r GinResponse) Redirect(code int, location string) error {
	r.c.Redirect(code, location)
	return nil
}

func (r GinResponse) File(path string) error {
	r.c.File(path)
	return nil
}
