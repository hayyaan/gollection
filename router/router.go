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
		GET(string, func(Request, Response) error)
		POST(string, func(Request, Response) error)
		DELETE(string, func(Request, Response) error)
		PATCH(string, func(Request, Response) error)
		PUT(string, func(Request, Response) error)
		OPTIONS(string, func(Request, Response) error)
		HEAD(string, func(Request, Response) error)

		// net/http
		HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
		Handle(pattern string, handler http.Handler)
	}

	Route interface {
		Method() string
		Path() string
		Handler() func(Request, Response) error
	}

	Request interface {
		Param(string) string
		Query(string) string
	}

	Response interface {
		AbortWithStatus(int) error
		JSON(int, interface{}) error
		XML(int, interface{}) error
		YAML(int, interface{}) error
		String(int, string, ...interface{}) error
		Redirect(int, string) error
		File(string) error
	}
)
