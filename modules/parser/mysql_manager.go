package parser

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"github.com/jerson/code-generator/modules/parser/platforms/mysql"
)

// MySQLManager ...
type MySQLManager struct {
	platform *mysql.Platform
}

// NewMySQLManager ...
func NewMySQLManager(ctx context.Base, driver, source string) (*MySQLManager, error) {
	platform, err := mysql.NewPlatform(ctx, driver, source)
	if err != nil {
		return nil, err
	}
	return &MySQLManager{platform: platform}, nil
}

//Views ...
func (m MySQLManager) Views() ([]models.View, error) {

	views, err := m.platform.Views("")
	if err != nil {
		return nil, err
	}

	var items []models.View
	for _, view := range views {
		items = append(items, models.View{
			Name: view.TableName,
			SQL:  view.ViewDefinition,
		})
	}

	return items, nil
}
