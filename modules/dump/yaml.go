package dump

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"gopkg.in/yaml.v2"
)

// Yaml ...
type Yaml struct {
	ctx    context.Base
	schema models.Schema
}

// NewYaml ...
func NewYaml(ctx context.Base, schema models.Schema) Yaml {
	return Yaml{ctx: ctx, schema: schema}
}

//Dump ...
func (y Yaml) Dump() (string, error) {
	data, err := yaml.Marshal(y.schema)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
