package dump

import (
	"encoding/json"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
)

// JSONDump ...
type JSONDump struct {
	ctx    context.Base
	schema models.Schema
}

// NewJSONDump ...
func NewJSONDump(ctx context.Base, schema models.Schema) JSONDump {
	return JSONDump{ctx: ctx, schema: schema}
}

//Dump ...
func (y JSONDump) Dump() (string, error) {
	data, err := json.Marshal(y.schema)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
