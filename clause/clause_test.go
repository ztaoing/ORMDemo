/**
* @Author:zhoutao
* @Date:2021/1/17 上午10:46
* @Desc:
 */

package clause

import (
	"reflect"
	"testing"
)

func testClause_Select(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Age DESC")
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age DESC LIMIT ?" {
		t.Fatal("failed to build sql")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build sql")
	}

}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testClause_Select(t)
	})
}
