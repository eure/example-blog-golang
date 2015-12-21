package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // load mysql driver for xorm
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

// UseEngine reutrns db resource
func UseEngine() *xorm.Engine {
	if engine == nil {
		engine = initEngine()
	}
	return engine
}

func initEngine() *xorm.Engine {
	dbUser := "user"
	dbPass := "pass"
	dbHost := "127.0.0.1"
	dbPort := 3306
	dbName := "blogdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	e, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// as default mapper, underscore is inserted on every capital case for table or field name.
	// (e.g.) UserID => user_i_d, IPAddress => i_p_address
	// to prevent this rule, use GonicMapper.
	e.SetMapper(core.NewCacheMapper(new(core.GonicMapper)))

	err = e.Ping()
	if err != nil {
		panic(err)
	}
	return e
}
