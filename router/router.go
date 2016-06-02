package router

import "net/http"

type (
	Engine interface {
		Use()
		Debug(bool)
		Run(string) error
		Router() Router
	}

	Router interface {
		GET(string, func(Response, Request) error)
		POST(string, func(Response, Request) error)
		DELETE(string, func(Response, Request) error)
		PATCH(string, func(Response, Request) error)
		PUT(string, func(Response, Request) error)
		OPTIONS(string, func(Response, Request) error)
		HEAD(string, func(Response, Request) error)

		// net/http
		HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
		Handle(pattern string, handler http.Handler)
	}

	Route interface {
		Method() string
		Path() string
		Handler() func(Response, Request) error
	}

	Request interface {
		Param(string) string
		Query(string) string
	}

	Response interface {
		AbortWithStatus(int) error
		String(int, string, ...interface{}) error
		HTML(int, string, interface{}) error
		JSON(int, interface{}) error
		XML(int, interface{}) error
		YAML(int, interface{}) error
		Redirect(int, string) error
		File(string) error
	}
)
