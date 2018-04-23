package mysql

import (
	"fmt"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/db"
	"github.com/jinzhu/gorm"
)

// MySQL ...
type MySQL struct {
	ctx context.Base
	cn  *gorm.DB
}

// NewPlatform ...
func NewPlatform(ctx context.Base, driver, source string) (*MySQL, error) {
	cn, err := db.Setup(ctx, driver, source)
	if err != nil {
		return nil, err
	}
	ctx.Set("DB", cn)
	return &MySQL{ctx: ctx, cn: cn}, nil
}

//Columns ...
func (m MySQL) Columns(table, database string) ([]Column, error) {

	if database == "" {
		database = "DATABASE()"
	} else {
		database = fmt.Sprintf("'%s'", database)
	}
	var result []Column
	err := m.cn.Raw(fmt.Sprintf(`SELECT COLUMN_NAME AS Field, COLUMN_TYPE AS Type, IS_NULLABLE AS 'Null', 
               COLUMN_KEY AS 'Key', COLUMN_DEFAULT AS 'Default', EXTRA AS Extra, COLUMN_COMMENT AS Comment, 
               CHARACTER_SET_NAME AS CharacterSet, COLLATION_NAME AS Collation 
               FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = %s AND TABLE_NAME = ?`, database), table).Scan(&result).Error

	return result, err
}

//Database ...
func (m MySQL) Database() (string, error) {
	var result string
	err := m.cn.Raw(`SELECT DATABASE()`).Row().Scan(&result)
	return result, err
}

//Indexes ...
func (m MySQL) Indexes(table, database string) ([]Index, error) {

	var result []Index
	if database != "" {
		err := m.cn.Raw(`SELECT TABLE_NAME AS 'Table', NON_UNIQUE AS Non_unique, INDEX_NAME AS Key_name,
                  SEQ_IN_INDEX AS Seq_in_index, COLUMN_NAME AS Column_name, COLLATION AS Collation,
                  CARDINALITY AS Cardinality, SUB_PART AS Sub_part, PACKED AS Packed,
                  NULLABLE AS 'Null', INDEX_TYPE AS Index_type, COMMENT AS Comment
                  FROM information_schema.STATISTICS WHERE TABLE_NAME = ? AND TABLE_SCHEMA =?`, table, database).Scan(&result).Error
		return result, err
	}
	err := m.cn.Raw(fmt.Sprintf(`SHOW INDEX FROM %s`, table)).Scan(&result).Error
	return result, err
}

//Views ...
func (m MySQL) Views(database string) ([]View, error) {

	var result []View
	err := m.cn.Raw(`SELECT * FROM information_schema.VIEWS WHERE TABLE_SCHEMA = ?`, database).Scan(&result).Error

	return result, err
}

//ForeignKeys ...
func (m MySQL) ForeignKeys(table, database string) ([]ForeignKey, error) {

	if database == "" {
		database = "DATABASE()"
	} else {
		database = fmt.Sprintf("'%s'", database)
	}

	var result []ForeignKey
	err := m.cn.Raw(fmt.Sprintf(`SELECT DISTINCT k.CONSTRAINT_NAME, k.COLUMN_NAME, k.REFERENCED_TABLE_NAME, 
               k.REFERENCED_COLUMN_NAME /*!50116 , c.update_rule, c.delete_rule */ 
               FROM information_schema.key_column_usage k /*!50116 
               INNER JOIN information_schema.referential_constraints c ON 
                 c.constraint_name = k.constraint_name AND 
                 c.table_name = '%s' */ WHERE k.table_name = '%s'
         		AND k.table_schema = %s /*!50116 AND c.constraint_schema = %s */
         		AND k.REFERENCED_COLUMN_NAME is not NULL`, table, table, database, database)).Scan(&result).Error
	return result, err

}

//Tables ...
func (m MySQL) Tables() ([]Table, error) {

	var result []Table
	rows, err := m.cn.Raw(`SHOW FULL TABLES WHERE Table_type = 'BASE TABLE'`).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var name string
		var tableType string

		rows.Scan(&name, &tableType)
		result = append(result, Table{
			Name:      name,
			TableType: tableType,
		})
	}

	return result, err

}
