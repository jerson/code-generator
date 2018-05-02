package parser

import (
	"encoding/json"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/keys"
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
	ctx              context.Base
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
	regexSpecialType := regexp.MustCompile("\\[type:(?P<type>[a-zA-Z0-9\\s]+)(, ?options:(?P<options>[\\w\\W}{]+))?]")
	return &MySQLManager{ctx: ctx, platform: platform, regexType: regexType, regexSpecialType: regexSpecialType}, nil
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
		Name:       base,
		Config:     *schemaConfig,
		Views:      views,
		Tables:     tables,
		Sequences:  sequences,
		Namespaces: namespaces,
	}, nil
}

//Database ...
func (m MySQLManager) Database() (string, error) {
	name, err := m.platform.Database()
	if err != nil {
		return "", err
	}
	return name, nil
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
			Name:       result.Name,
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
			Name: result.TableName,
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
		columnSpecialType, err := m.parseSpecialType(result.Comment)
		if err != nil {
			return nil, err
		}

		key := keys.Unknown

		switch result.Key {
		case "PRI":
			key = keys.Primary
			break
		case "MUL":
			key = keys.Multiple
			break
		}

		items = append(items, models.Column{
			Name:          result.Field,
			SpecialType:   columnSpecialType,
			Type:          columnType.Value,
			Length:        columnType.Length,
			Precision:     columnType.Precision,
			Scale:         columnType.Scale,
			Fixed:         columnType.Fixed,
			Unsigned:      strings.Contains(result.Type, "unsigned"),
			NotNull:       result.Null != "YES",
			Default:       result.Default,
			AutoIncrement: strings.Contains(result.Extra, "auto_increment"),
			Key:           key,
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

	typeString := strings.ToLower(strings.TrimSpace(results["type"]))
	lengthString := strings.Trim(strings.Trim(results["length"], ")"), "(")
	if lengthString != "" && !strings.Contains(lengthString, ",") {
		length, err := strconv.Atoi(lengthString)
		if err != nil {
			return nil, err
		}
		columnType.Length = length
	}

	value := types.Unknown

	switch typeString {
	case "int", "smallint", "mediumint":
		value = types.Integer
	case "tinyint":
		if columnType.Length == 1 {
			value = types.Boolean
		} else {
			value = types.SmallInt
		}
	case "bigint":
		value = types.BigInt
	case "decimal":
		value = types.Decimal
	case "double", "float", "real":
		value = types.Float
	case "boolean", "bit", "serial":
		value = types.Boolean
	case "date":
		value = types.Date
	case "datetime":
		value = types.Datetime
	case "time":
		value = types.Time
	case "year":
		value = types.Year
	case "timestamp":
		value = types.Datetime
	case "char":
		columnType.Fixed = true
		value = types.String
	case "varchar":
		value = types.String
	case "tinytext", "text", "mediumtext", "longtext":
		value = types.Text
	case "binary":
		columnType.Fixed = true
		value = types.Binary
	case "varbinary":
		value = types.Binary
	case "tinyblob", "mediumblob", "blob", "longblob":
		value = types.Blob
	case "enum":
		value = types.String
		columnType.Options = &models.TypeExtraOptions{Values: strings.Split(lengthString, ",")}
	case "json":
		value = types.JSON
	}
	columnType.Value = value

	return columnType, nil
}

//parseSpecialType ...
func (m MySQLManager) parseSpecialType(comment string) (*models.Type, error) {

	log := m.ctx.GetLogger("parseSpecialType")
	results := util.GetRegexParams(m.regexSpecialType, comment)
	columnType := &models.Type{}

	typeString := strings.ToLower(results["type"])
	optionsString := strings.TrimSpace(results["options"])

	if typeString == "" {
		return nil, nil
	}

	if optionsString != "" {
		var options *models.TypeExtraOptions
		err := json.Unmarshal([]byte(optionsString), &options)
		if err != nil {
			log.Warn(err)
		}

		if options != nil {
			columnType.Options = options
			columnType.Length = options.Length
		}
	}

	value := types.String
	switch typeString {
	case "email":
		value = types.SpecialEmail
		break
	case "password":
		value = types.SpecialPassword
		break
	case "url":
		value = types.SpecialURL
		break
	case "html":
		value = types.SpecialHTML
		break
	default:
		value = types.Value(typeString)
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

	list := map[string]*models.Index{}

	for _, value := range results {

		if list[value.KeyName] == nil {
			var flags []string

			if value.IndexType == "FULLTEXT" {
				flags = append(flags, "FULLTEXT")
			} else if value.IndexType == "SPATIAL" {
				flags = append(flags, "SPATIAL")
			}

			list[value.KeyName] = &models.Index{
				Name:      value.KeyName,
				IsUnique:  value.NonUnique == "0",
				IsPrimary: value.KeyName == "PRIMARY",
				Flags:     flags,
			}
		}
		list[value.KeyName].Columns = append(list[value.KeyName].Columns, models.Identifier{Name: value.ColumnName})

	}
	var items []models.Index
	for _, result := range list {
		items = append(items, *result)
	}

	return items, nil
}

//ForeignKeys ...
func (m MySQLManager) ForeignKeys(table string) ([]models.ForeignKey, error) {

	results, err := m.platform.ForeignKeys(table, "")
	if err != nil {
		return nil, err
	}

	list := map[string]*models.ForeignKey{}

	for _, value := range results {

		if list[value.ConstraintName] == nil {
			if value.DeleteRule == "RESTRICT" {
				value.DeleteRule = ""
			}
			if value.UpdateRule == "RESTRICT" {
				value.UpdateRule = ""
			}
			list[value.ConstraintName] = &models.ForeignKey{
				Name:               value.ConstraintName,
				LocalTable:         table,
				LocalColumnName:    []models.Identifier{},
				ForeignColumnNames: []models.Identifier{},
				ForeignTableName: models.Identifier{
					Name: value.ReferencedTableName,
				},
				Options: models.ForeignKeyOptions{
					OnUpdate: value.UpdateRule,
					OnDelete: value.DeleteRule,
				},
			}
		}
		list[value.ConstraintName].LocalColumnName = append(list[value.ConstraintName].LocalColumnName, models.Identifier{Name: value.ColumnName})
		list[value.ConstraintName].ForeignColumnNames = append(list[value.ConstraintName].ForeignColumnNames, models.Identifier{Name: value.ReferencedColumnName})

	}

	var items []models.ForeignKey

	for _, result := range list {
		items = append(items, *result)
	}

	return items, nil
}
