/**
* @Author:zhoutao
* @Date:2021/1/14 下午12:46
* @Desc: engine 是ORMDemo与用户交互的入口
 */

package ORMDemo

import (
	"database/sql"
	"github.com/ztaoing/ORMDemo/log"
	"github.com/ztaoing/ORMDemo/session"
)

type Engine struct {
	db *sql.DB
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("failed to close database")
	}
	log.Info("close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db)
}

func NewEngine(driver, source string) (e *Engine, err error) {
	//连接数据库
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//检查数据库连接是否正常
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{
		db: db,
	}
	log.Info("connect database success")
	return
}
