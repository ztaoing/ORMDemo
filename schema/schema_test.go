/**
* @Author:zhoutao
* @Date:2021/1/14 下午2:56
* @Desc:
 */

package schema

import (
	"github.com/ztaoing/ORMDemo/dialect"
	"testing"
)

type User struct {
	Name string `ormdemo:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
