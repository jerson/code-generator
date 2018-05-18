package mapper

import (
	"fmt"
	"github.com/jerson/code-generator/modules/context"
	"reflect"
	"strings"
)

// GORM ...
type GORM struct {
	ctx context.Base
}

// NewGORM ...
func NewGORM(ctx context.Base) *GORM {
	return &GORM{}
}

// Scan ...
func (g GORM) Scan(entities ...interface{}) {

	for _, entity := range entities {
		g.readTags(entity)
	}

}
func (g GORM) readTags(entity interface{}) {

	v := reflect.ValueOf(entity)

	for i := 0; i < v.NumField(); i++ {
		//fmt.Println(v.Type().Field(i))
		tag := v.Type().Field(i).Tag.Get("gorm")

		if tag == "" || tag == "-" {
			continue
		}

		args := strings.Split(tag, ",")

		fmt.Println(args)
	}

}
