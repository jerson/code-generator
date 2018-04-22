package context

import (
	"github.com/sirupsen/logrus"
)

// Base ...
type Base interface {
	Set(name string, value interface{})
	Get(name string) interface{}
	Close()
	GetToken() string
	GetLogger(tag string) *logrus.Entry
}
