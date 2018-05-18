package mapper

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/mapper/samples"
	"testing"
)

func TestGORM(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()
	cn := NewGORM(ctx)
	cn.Scan(samples.Category{}, samples.Post{})
}
