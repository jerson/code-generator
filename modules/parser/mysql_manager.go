package parser

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"github.com/jerson/code-generator/modules/parser/platforms/mysql"
	"github.com/jerson/code-generator/modules/parser/types"
	"github.com/jerson/code-generator/modules/util"
	"regexp"
	"strconv"
	"strings"
)

// MySQLManager ...
type MySQLManager struct {
	platform         *mysql.Platform
	regexType        *regexp.Regexp
	regexSpecialType *regexp.Regexp
}

// NewMySQLManager ...
func NewMySQLManager(ctx context.Base, driver, source string) (*MySQLManager, error) {
	platform, err := mysql.NewPlatform(ctx, driver, source)
	if err != nil {
		return nil, err
	}

	regexType := regexp.MustCompile("^(?P<type>[A-Za-z0-9]+)(?P<length>\\([0-9A-Za-z\\'\\'\\,]+\\))?$")
	regexSpecialType := regexp.MustCompile("^(?P<type>[A-Za-z0-9]+)(?P<length>\\([0-9A-Za-z\\'\\'\\,]+\\))?$")
	return &MySQLManager{platform: platform, regexType: regexType, regexSpecialType: regexSpecialType}, nil
}

//Schema ...
func (m MySQLManager) Schema() (*models.Schema, error) {
	base, err := m.Database()
	if err != nil {
		return nil, err
	}
	views, err := m.Views()
	if err != nil {
		return nil, err
	}
	tables, err := m.Tables()
	if err != nil {
		return nil, err
	}
	sequences, err := m.Sequences()
	if err != nil {
		return nil, err
	}
	namespaces, err := m.Namespaces()
	if err != nil {
		return nil, err
	}
	schemaConfig, err := m.SchemaConfig()
	if err != nil {
		return nil, err
	}
	return &models.Schema{
		Base:       *base,
		Config:     *schemaConfig,
		Views:      views,
		Tables:     tables,
		Sequences:  sequences,
		Namespaces: namespaces,
	}, nil
}

//Database ...
func (m MySQLManager) Database() (*models.Base, error) {
	name, err := m.platform.Database()
	if err != nil {
		return nil, err
	}
	return &models.Base{Name: name}, nil
}

//SchemaConfig ...
func (m MySQLManager) SchemaConfig() (*models.SchemaConfig, error) {
	return &models.SchemaConfig{}, nil
}

//Sequences ...
func (m MySQLManager) Sequences() ([]models.Sequence, error) {
	return []models.Sequence{}, nil
}

//Namespaces ...
func (m MySQLManager) Namespaces() ([]string, error) {
	return []string{}, nil
}

//Tables ...
func (m MySQLManager) Tables() ([]models.Table, error) {
	results, err := m.platform.Tables()
	if err != nil {
		return nil, err
	}

	var items []models.Table
	for _, result := range results {

		indexes, err := m.Indexes(result.Name)
		if err != nil {
			return nil, err
		}
		foreignKeys, err := m.ForeignKeys(result.Name)
		if err != nil {
			return nil, err
		}
		columns, err := m.Columns(result.Name)
		if err != nil {
			return nil, err
		}

		items = append(items, models.Table{
			Base:       models.Base{Name: result.Name},
			Columns:    columns,
			Indexes:    indexes,
			ForeignKey: foreignKeys,
		})
	}

	return items, nil
}

//Views ...
func (m MySQLManager) Views() ([]models.View, error) {

	results, err := m.platform.Views("")
	if err != nil {
		return nil, err
	}

	var items []models.View
	for _, result := range results {
		items = append(items, models.View{
			Base: models.Base{Name: result.TableName},
			SQL:  result.ViewDefinition,
		})
	}

	return items, nil
}

//Columns ...
func (m MySQLManager) Columns(table string) ([]models.Column, error) {

	results, err := m.platform.Columns(table, "")
	if err != nil {
		return nil, err
	}

	var items []models.Column
	for _, result := range results {

		columnType, err := m.parseType(result.Type)
		if err != nil {
			return nil, err
		}
		columnSpecialTYpe, err := m.parseSpecialType(result.Comment)
		if err != nil {
			return nil, err
		}

		items = append(items, models.Column{
			Base:          models.Base{Name: result.Field},
			SpecialType:   columnSpecialTYpe,
			Type:          columnType.Value,
			Length:        columnType.Length,
			Precision:     columnType.Precision,
			Scale:         columnType.Scale,
			Fixed:         columnType.Fixed,
			Unsigned:      strings.Contains(result.Type, "unsigned"),
			NotNull:       result.Null != "YES",
			Default:       result.Default,
			AutoIncrement: strings.Contains(result.Extra, "auto_increment"),
			PlatformOptions: models.PlatformOptions{
				Collation:    result.Collation,
				CharacterSet: result.CharacterSet,
			},
			Comment: result.Comment,
		})
	}

	return items, nil
}

//parseType ...
func (m MySQLManager) parseType(typeData string) (*models.Type, error) {

	results := util.GetRegexParams(m.regexType, typeData)

	columnType := &models.Type{}
	columnType.Name = typeData

	typeString := results["type"]
	lengthString := strings.Trim(strings.Trim(results["length"], ")"), "(")
	if lengthString != "" {
		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return nil, err
		}
		columnType.Length = length
	}

	value := types.String
	switch strings.ToLower(typeString) {
	case "tinyint":
		if columnType.Length == 1 {
			value = types.Boolean
		} else {
			value = types.SmallInt
		}
		break
	case "bigint":
		value = types.BigInt
		break
	case "int":
		value = types.Integer
		break
	case "char":
		columnType.Fixed = true
		value = types.String
		break
	case "double":
	case "float":
	case "real":
		value = types.Float
		break
	case "text":
		value = types.Text
		break
	case "bool":
		value = types.Boolean
		break
	case "date":
		value = types.Date
		break
	case "datetime":
		value = types.Datetime
		break
	case "time":
		value = types.Time
		break
	case "timestamp":
		value = types.Datetime
		break
	default:
		value = types.Unknown
		break
	}
	columnType.Value = value

	return columnType, nil
}

//parseSpecialType ...
func (m MySQLManager) parseSpecialType(comment string) (*models.Type, error) {

	results := util.GetRegexParams(m.regexSpecialType, comment)

	columnType := &models.Type{}

	typeString := results["type"]
	lengthString := strings.Trim(strings.Trim(results["length"], ")"), "(")
	if lengthString != "" {
		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return nil, err
		}
		columnType.Length = length
	}

	value := types.String
	switch strings.ToLower(typeString) {
	case "tinyint":
		if columnType.Length == 1 {
			value = types.Boolean
		} else {
			value = types.SmallInt
		}
		break
	default:
		value = types.Unknown
		break
	}
	columnType.Value = value

	return columnType, nil
}

//Indexes ...
func (m MySQLManager) Indexes(table string) ([]models.Index, error) {

	results, err := m.platform.Indexes(table, "")
	if err != nil {
		return nil, err
	}

	var items []models.Index
	for _, result := range results {

		var flags []string

		if result.IndexType == "FULLTEXT" {
			flags = append(flags, "FULLTEXT")
		} else if result.IndexType == "SPATIAL" {
			flags = append(flags, "SPATIAL")
		}

		var columns []models.Identifier
		identifier := models.Identifier{Base: models.Base{Name: result.ColumnName}}
		columns = append(columns, identifier)

		items = append(items, models.Index{
			Base:      models.Base{Name: result.KeyName},
			Columns:   columns,
			IsUnique:  result.NonUnique == "0",
			IsPrimary: result.KeyName == "PRIMARY",
			Flags:     flags,
		})
	}

	return items, nil
}

//ForeignKeys ...
func (m MySQLManager) ForeignKeys(table string) ([]models.ForeignKey, error) {

	results, err := m.platform.ForeignKeys(table, "")
	if err != nil {
		return nil, err
	}

	list := map[string]*foreignKeyTemp{}

	for _, value := range results {

		if list[value.ConstraintName] == nil {
			if value.DeleteRule == "RESTRICT" {
				value.DeleteRule = ""
			}
			if value.UpdateRule == "RESTRICT" {
				value.UpdateRule = ""
			}
			list[value.ConstraintName] = &foreignKeyTemp{
				Name:         value.ConstraintName,
				Local:        []models.Identifier{},
				Foreign:      []models.Identifier{},
				ForeignTable: value.ReferencedTableName,
				OnDelete:     value.DeleteRule,
				OnUpdate:     value.UpdateRule,
			}
		}
		if list[value.ConstraintName] != nil {
			list[value.ConstraintName].Local = append(list[value.ConstraintName].Local, models.Identifier{Base: models.Base{Name: value.ColumnName}})
			list[value.ConstraintName].Foreign = append(list[value.ConstraintName].Foreign, models.Identifier{Base: models.Base{Name: value.ReferencedColumnName}})
		}

	}

	var items []models.ForeignKey

	for _, result := range list {

		items = append(items, models.ForeignKey{
			Base:               models.Base{Name: result.Name},
			LocalTable:         table,
			LocalColumnName:    result.Local,
			ForeignColumnNames: result.Foreign,
			ForeignTableName: models.Identifier{
				Base: models.Base{Name: result.ForeignTable},
			},
			Options: models.ForeignKeyOptions{
				OnUpdate: result.OnUpdate,
				OnDelete: result.OnDelete,
			},
		})

	}

	return items, nil
}

type foreignKeyTemp struct {
	Name         string
	ForeignTable string
	Local        []models.Identifier
	Foreign      []models.Identifier
	OnDelete     string
	OnUpdate     string
}
