/**
* @Author:zhoutao
* @Date:2021/1/14 下午2:12
* @Desc: 对象和表的转换，给定一个任意的对象，转换为关系型数据库中的表结构
 */

package schema

import (
	"github.com/ztaoing/ORMDemo/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string //约束条件:如非空、主键等，go语言通过tag实现，Java和Python等通过注解实现
}

type Schema struct {
	Model      interface{}       //对象
	Name       string            //表名
	Fields     []*Field          //字段
	FieldNames []string          //包含所有的字段名
	filedMap   map[string]*Field //方便直接使用
}

func (s *Schema) GetField(name string) *Field {
	return s.filedMap[name]
}

//将任意的对象解析为schema
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		filedMap: make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		f := modelType.Field(i)
		//非匿名的+可导出的
		if !f.Anonymous && ast.IsExported(f.Name) {
			field := &Field{
				Name: f.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(f.Type))),
			}
			if v, ok := f.Tag.Lookup("ormdemo"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, f.Name)
			schema.filedMap[f.Name] = field
		}
	}
	return schema
}
