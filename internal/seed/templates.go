package seed

import (
	"embed"
	"fmt"
	"io/fs"
	"strings"
)

const (
	// templateNameNumParts represents the number of parts in a template name
	// names are all formatted as datum.<name>.json
	templateNameNumParts = 3
)

//go:embed templates/*/*.json
var jsonSchemaFS embed.FS

var (
	templateDirectory = "templates/jsonschemas"
)

// Template represents a template to be seeded
type Template struct {
	// Name is the name of the template
	Name string
	// JSONConfig is the JSONschema configuration for the template
	JSONConfig []byte
}

// getTemplates gets all the templates from the jsonschema directory
func getTemplates(dir string) (templates []Template, err error) {
	err = fs.WalkDir(jsonSchemaFS, dir, func(path string, d fs.DirEntry, walkErr error) error {
		if d == nil {
			return walkErr
		}

		if d.IsDir() {
			return nil
		}

		schema, err := jsonSchemaFS.ReadFile(path)
		if err != nil {
			return err
		}

		nameSplit := strings.Split(d.Name(), ".")

		if len(nameSplit) != templateNameNumParts {
			return fmt.Errorf("%w: %s", ErrInvalidTemplateName, d.Name())
		}

		templates = append(templates, Template{
			Name:       nameSplit[1],
			JSONConfig: schema,
		})

		return nil
	})

	return
}
