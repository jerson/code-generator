package dump

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"gopkg.in/yaml.v2"
)

// YamlDump ...
type YamlDump struct {
	ctx    context.Base
	schema models.Schema
}

// NewYamlDump ...
func NewYamlDump(ctx context.Base, schema models.Schema) YamlDump {
	return YamlDump{ctx: ctx, schema: schema}
}

//Dump ...
func (y YamlDump) Dump() (string, error) {
	data, err := yaml.Marshal(y.schema)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
