package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinEngine struct {
	gin *gin.Engine
}

func NewGin() GinEngine {
	gin.SetMode(gin.ReleaseMode)
	return GinEngine{gin: gin.New()}
}

func (r GinEngine) Debug(d bool) {
	if d {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}

// Engine implementation

func (r GinEngine) Use() {
}

func (r GinEngine) Router() Router {
	return r
}

func (r GinEngine) Run(addr string) error {
	return r.gin.Run(addr)
}

// Router implementation

func (r GinEngine) GET(path string, h func(Request, Response) error) {
	r.gin.GET(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) POST(path string, h func(Request, Response) error) {
	r.gin.POST(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) DELETE(path string, h func(Request, Response) error) {
	r.gin.DELETE(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) PATCH(path string, h func(Request, Response) error) {
	r.gin.PATCH(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) PUT(path string, h func(Request, Response) error) {
	r.gin.PUT(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) OPTIONS(path string, h func(Request, Response) error) {
	r.gin.OPTIONS(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) HEAD(path string, h func(Request, Response) error) {
	r.gin.HEAD(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) Handle(pattern string, handler http.Handler) {
	panic("gin Handle not yet implemented") // TODO
}

func (r GinEngine) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
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

func (r GinResponse) String(code int, s string, v ...interface{}) error {
	r.c.String(code, s, v)
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
