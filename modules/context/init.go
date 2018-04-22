package context

import (
	"github.com/jerson/code-generator/modules/config"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

// template ...
type template struct {
	Token    string
	Category string
	Logger   *logrus.Entry
}

// GetToken ...
func (r *template) GetToken() string {
	return r.Token
}

// GetLogger ...
func (r *template) GetLogger(tag string) *logrus.Entry {
	if r.Logger != nil {
		return r.Logger.WithFields(map[string]interface{}{
			"tag": tag,
		})
	}

	log := logrus.New()

	if config.Vars.Debug {
		log.SetLevel(logrus.DebugLevel)
	}
	r.Logger = log.WithFields(map[string]interface{}{
		"category": r.Category,
		"tag":      tag,
		"token":    r.Token,
	})

	return r.Logger
}

// GetCacheResponse ...
func (r template) GetCacheResponse(key string) ([]byte, error) {
	return nil, errors.New("not implemented")
}
