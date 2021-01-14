/**
* @Author:zhoutao
* @Date:2021/1/14 下午1:05
* @Desc:
 */

package session

import (
	"database/sql"
	"os"
	"testing"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)

}

func NewSessionT() *Session {
	return NewSession(TestDB)
}
func TestSession_Exec(t *testing.T) {
	s := NewSessionT()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values(?),(?)", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2,but got :", count)
	}
}

func TestNewSession_Query(t *testing.T) {
	s := NewSessionT()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
