package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPIHandler(router *Router) (err error) {
	path := "/api-docs"
	method := http.MethodGet
	name := "APIDocs"

	route := echo.Route{
		Name:        name,
		Method:      method,
		Path:        path,
		Middlewares: mw,
		Handler: echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, router.OAS)
		}),
	}

	if err := router.AddEchoOnlyRoute(path, method, route); err != nil {
		return err
	}

	return nil
}
