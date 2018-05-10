package dump

import (
	"encoding/json"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
)

// JSON ...
type JSON struct {
	ctx    context.Base
	schema models.Schema
}

// NewJSON ...
func NewJSON(ctx context.Base, schema models.Schema) JSON {
	return JSON{ctx: ctx, schema: schema}
}

//Dump ...
func (y JSON) Dump() (string, error) {
	data, err := json.Marshal(y.schema)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
