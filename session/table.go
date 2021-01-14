/**
* @Author:zhoutao
* @Date:2021/1/14 下午3:06
* @Desc: 操作数据库表相关的操作
 */

package session

import (
	"errors"
	"fmt"
	"github.com/ztaoing/ORMDemo/log"
	"github.com/ztaoing/ORMDemo/schema"
	"reflect"
	"strings"
)

// 用于给refTable赋值，解析操作是比较耗时的，所以讲解析的结果保存在refTable中
// 即使调用多次，如果传入的结构体名称不发生变化，则不会更新refTable
func (s *Session) Model(value interface{}) *Session {
	// nil or different model will update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

//返回refTable
func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model does not be set")
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	if table == nil {
		return errors.New("refTable is nil")
	}
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IS EXISTS %s", s.refTable.Name)).Exec()
	return err
}

//判断是否存在
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.refTable.Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name

}
