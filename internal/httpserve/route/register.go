package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerWorkspaceHandler registers the workspace handler and route
func registerWorkspaceHandler(router *Router) (err error) {
	path := "/workspace"
	method := http.MethodPost
	name := "Workspace"

	route := echo.Route{
		Name:   name,
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.WorkspaceHandler(c)
		},
	}

	registerOperation := router.Handler.BindWorkspaceHandler()

	if err := router.Addv1Route(path, method, registerOperation, route); err != nil {
		return err
	}

	return nil
}
