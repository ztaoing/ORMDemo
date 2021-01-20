/**
* @Author:zhoutao
* @Date:2021/1/20 下午2:01
* @Desc: 实现记录增删改查
 */

package session

import (
	"github.com/ztaoing/ORMDemo/clause"
	"reflect"
)

//将已经存在的对象的每一个字段的值平铺开
func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	// 通过set构造每一个字句
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}
	s.clause.Set(clause.VALUES, recordValues...)
	// 根据传入的顺序构造出最终的SQL语句
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	// 执行SQL
	res, err := s.Raw(sql, vars).Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//根据平铺的字段构造出对象
func (s *Session) Find(value interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(value))
	//获取切片的单个元素的类型
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}
	//遍历每一行结果
	for rows.Next() {
		//利用反射创建destType的实例dest
		dest := reflect.New(destType).Elem()
		var values []interface{}
		//将dest的所有字段平铺开，构造切片values
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		//调用rows.Scan将该行记录每一列的值，依次赋值给values中的每一个字段
		if err := rows.Scan(values...); err != nil {
			return err
		}
		//将dest添加到切片destSlice中
		destSlice.Set(reflect.Append(destSlice, dest))
	}

}
