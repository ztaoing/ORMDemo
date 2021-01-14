/**
* @Author:zhoutao
* @Date:2021/1/14 下午1:34
* @Desc: 将go语言的类型映射为数据库中类型，而且不同的数据库支持的数据类型也是有差异的。orm框架往往需要兼容多种数据库
 */

package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	//将go语言的类型转换为该数据库的数据类型
	DataTypeOf(typ reflect.Value) string
	//返回某个表是否存在的SQL语句
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
