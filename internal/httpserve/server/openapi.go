package server

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3gen"

	"github.com/datumforge/datum/pkg/rout"
)

// NewOpenAPISpec creates a new OpenAPI 3.1.0 specification based on the configured go interfaces and the operation types appended within the individual handlers
func NewOpenAPISpec() (*openapi3.T, error) {
	schemas := make(openapi3.Schemas)
	responses := make(openapi3.ResponseBodies)
	parameters := make(openapi3.ParametersMap)
	requestbodies := make(openapi3.RequestBodies)
	securityschemes := make(openapi3.SecuritySchemes)
	examples := make(openapi3.Examples)

	generator := openapi3gen.NewGenerator(openapi3gen.UseAllExportedFields())
	for key, val := range openAPISchemas {
		ref, err := generator.NewSchemaRefForValue(val, schemas)
		if err != nil {
			return nil, err
		}

		schemas[key] = ref
	}

	errorResponse := &openapi3.SchemaRef{
		Ref: "#/components/schemas/ErrorResponse",
	}

	_, err := openapi3gen.NewSchemaRefForValue(&rout.StatusError{}, schemas)
	if err != nil {
		return nil, err
	}

	internalServerError := openapi3.NewResponse().
		WithDescription("Internal Server Error").
		WithContent(openapi3.NewContentWithJSONSchemaRef(errorResponse))
	responses["InternalServerError"] = &openapi3.ResponseRef{Value: internalServerError}

	badRequest := openapi3.NewResponse().
		WithDescription("Bad Request").
		WithContent(openapi3.NewContentWithJSONSchemaRef(errorResponse))
	responses["BadRequest"] = &openapi3.ResponseRef{Value: badRequest}

	unauthorized := openapi3.NewResponse().
		WithDescription("Unauthorized").
		WithContent(openapi3.NewContentWithJSONSchemaRef(errorResponse))
	responses["Unauthorized"] = &openapi3.ResponseRef{Value: unauthorized}

	conflict := openapi3.NewResponse().
		WithDescription("Conflict").
		WithContent(openapi3.NewContentWithJSONSchemaRef(errorResponse))
	responses["Conflict"] = &openapi3.ResponseRef{Value: conflict}

	return &openapi3.T{
		OpenAPI: "3.1.0",
		Info: &openapi3.Info{
			Title:   "Datum Cloud OpenAPI 3.1.0 Specifications",
			Version: "v1.0.0",
			Contact: &openapi3.Contact{
				Name:  "Datum",
				Email: "support@datum.net",
				URL:   "https://datum.net",
			},
			License: &openapi3.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0",
			},
		},
		Paths: openapi3.NewPaths(),
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Datum Cloud API Server (local)",
				URL:         "http://localhost:17610/v1",
			},
		},

		Components: &openapi3.Components{
			Schemas:         schemas,
			Responses:       responses,
			Parameters:      parameters,
			RequestBodies:   requestbodies,
			SecuritySchemes: securityschemes,
			Examples:        examples,
		},
	}, nil
}

// openAPISchemas is a mapping of types to auto generate schemas for - these specifically live under the OAS "schema" type so that we can simply make schemaRef's to them and not have to define them all individually in the OAS paths
var openAPISchemas = map[string]any{
	"ErrorResponse": &rout.StatusError{},
}

// OAuth2 is a struct that represents an OAuth2 security scheme
type OAuth2 struct {
	AuthorizationURL string
	TokenURL         string
	RefreshURL       string
	Scopes           map[string]string
}

// Scheme returns the OAuth2 security scheme
func (i *OAuth2) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type: "oauth2",
		Flows: &openapi3.OAuthFlows{
			AuthorizationCode: &openapi3.OAuthFlow{
				AuthorizationURL: i.AuthorizationURL,
				TokenURL:         i.TokenURL,
				RefreshURL:       i.RefreshURL,
				Scopes:           i.Scopes,
			},
		},
	}
}

// OpenID is a struct that represents an OpenID Connect security scheme
type OpenID struct {
	ConnectURL string
}

// Scheme returns the OpenID Connect security scheme
func (i *OpenID) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type:             "openIdConnect",
		OpenIdConnectUrl: i.ConnectURL,
	}
}

// APIKey is a struct that represents an API Key security scheme
type APIKey struct {
	Name string
}

// Scheme returns the API Key security scheme
func (k *APIKey) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type: "http",
		In:   "header",
		Name: k.Name,
	}
}

// Basic is a struct that represents a Basic Auth security scheme
type Basic struct {
	Username string
	Password string
}

// Scheme returns the Basic Auth security scheme
func (b *Basic) Scheme() *openapi3.SecurityScheme {
	return &openapi3.SecurityScheme{
		Type:   "http",
		Scheme: "basic",
	}
}
