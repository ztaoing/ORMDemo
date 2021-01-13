/**
* @Author:zhoutao
* @Date:2021/1/13 下午4:10
* @Desc:用于实现与数据库的交互
 */

package session

import (
	"database/sql"
	"github.com/ztaoing/ORMDemo/log"
	"strings"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder //拼接SQL语句
	sqlVars []interface{}   //SQL语句中占位符的对应值
}

func NewSession(db *sql.DB) *Session {
	return &Session{
		db: db,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil

}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values)
	return s
}

func (s *Session) Exec() (res sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if res, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars); err != nil {
		log.Error(err)
	}
	return
}
