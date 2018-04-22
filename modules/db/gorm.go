package db

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// Setup ...
func Setup(ctx context.Base, driver, source string) (*gorm.DB, error) {

	cn, err := gorm.Open(driver, source)
	if err != nil {
		return cn, err
	}

	//log := ctx.GetLogger("DB.gorm")
	//cn.SetLogger(Logger{Instance: log})
	//cn.LogMode(config.Vars.Debug)
	//cn.LogMode(false)

	return cn, nil
}

// Logger ...
type Logger struct {
	Instance *logrus.Entry
}

// Print ...
func (logger Logger) Print(values ...interface{}) {
	logger.Instance.Error(values)
}
