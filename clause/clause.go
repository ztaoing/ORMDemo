/**
* @Author:zhoutao
* @Date:2021/1/14 下午3:57
* @Desc:
 */

package clause

import "strings"

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

//根据Type调用对应的generator,生成该字句对应的SQL语句
func (c *Clause) Set(typ Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.sqlVars = make(map[Type][]interface{})
	}
	//根据类型获取对应类型的处理函数
	sql, vars := generators[typ](vars...)
	//赋值语句和值
	c.sql[typ] = sql
	c.sqlVars[typ] = vars
}

//根据传入的Type的顺序，构造出最终的SQL语句
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	//按传入的顺序
	for _, order := range orders {
		if v, ok := c.sql[order]; ok {
			sqls = append(sqls, v)
			vars = append(vars, c.sqlVars[order]...)
		}
	}
	// WHERE  FROM
	return strings.Join(sqls, " "), vars
}
