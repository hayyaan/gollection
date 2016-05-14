package router

import (
	"github.com/gin-gonic/gin"
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

func (r GinEngine) GET(path string, h func(Request, Response)) {
	r.gin.GET(path, func(c *gin.Context) {
		h(GinRequest{c}, GinResponse{c})
	})
}

func (r GinEngine) POST(path string, h func(Request, Response)) {
	r.gin.POST(path)
}

func (r GinEngine) DELETE(path string, h func(Request, Response)) {
	r.gin.DELETE(path)
}

func (r GinEngine) PATCH(path string, h func(Request, Response)) {
	r.gin.PATCH(path)
}

func (r GinEngine) PUT(path string, h func(Request, Response)) {
	r.gin.PUT(path)
}

func (r GinEngine) OPTIONS(path string, h func(Request, Response)) {
	r.gin.OPTIONS(path)
}

func (r GinEngine) HEAD(path string, h func(Request, Response)) {
	r.gin.HEAD(path)
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

func (r GinResponse) AbortWithStatus(code int) {
	r.c.AbortWithStatus(code)
}

func (r GinResponse) JSON(code int, obj interface{}) {
	r.c.JSON(code, obj)
}

func (r GinResponse) XML(code int, obj interface{}) {
	r.c.XML(code, obj)
}

func (r GinResponse) YAML(code int, obj interface{}) {
	r.c.YAML(code, obj)
}

func (r GinResponse) String(code int, s string, v ...interface{}) {
	r.c.String(code, s, v)
}

func (r GinResponse) Redirect(code int, location string) {
	r.c.Redirect(code, location)
}

func (r GinResponse) File(path string) {
	r.c.File(path)
}
