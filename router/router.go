package router

type (
	Engine interface {
		Use()
		Debug(bool)
		Run(string) error
		Router() Router
	}

	Router interface {
		GET(string, func(Request, Response))
		POST(string, func(Request, Response))
		DELETE(string, func(Request, Response))
		PATCH(string, func(Request, Response))
		PUT(string, func(Request, Response))
		OPTIONS(string, func(Request, Response))
		HEAD(string, func(Request, Response))
	}

	Route interface {
		Method() string
		Path() string
		Handler() func(Request, Response)
	}

	Request interface {
		Param(string) string
		Query(string) string
	}

	Response interface {
		AbortWithStatus(int)
		JSON(int, interface{})
		XML(int, interface{})
		YAML(int, interface{})
		String(int, string, ...interface{})
		Redirect(int, string)
		File(string)
	}
)
