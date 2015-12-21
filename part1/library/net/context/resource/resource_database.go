package resource

import (
	"github.com/go-xorm/xorm"
	"golang.org/x/net/context"

	"github.com/eure/example-blog-golang/library/database"
)

// GetDB returns database engine
func GetDB(c context.Context) *xorm.Engine {
	v, ok := c.Value("database").(*xorm.Engine)
	if !ok {
		return nil
	}
	return v
}

// UseDB returns database engine
func UseDB(c context.Context) *xorm.Engine {
	db := GetDB(c)
	if db == nil {
		db = database.UseEngine()
		c = context.WithValue(c, "database", db)
	}
	return db
}

// UseTx returns database transaction
func UseTx(c context.Context) *xorm.Session {
	tx := c.Value("transaction").(*xorm.Session)
	if tx == nil {
		db := UseDB(c)
		tx := db.NewSession()
		tx.Begin()
		c = context.WithValue(c, "transaction", tx)
	}
	return tx
}

// Commit commits database transaction
func Commit(c context.Context) error {
	tx := c.Value("transaction").(*xorm.Session)
	if tx == nil {
		return nil
	}

	err := tx.Commit()
	if err != nil {
		return err
	}

	c = context.WithValue(c, "transaction", nil)
	tx.Close()
	return nil
}

// Rollback rollbacks database transaction
func Rollback(c context.Context) error {
	tx := c.Value("transaction").(*xorm.Session)
	if tx == nil {
		return nil
	}

	err := tx.Rollback()
	if err != nil {
		return err
	}

	c = context.WithValue(c, "transaction", nil)
	tx.Close()
	return nil
}

// Release releases all resources
func Release(c context.Context) {
	if c == nil {
		return
	}

	tx := c.Value("transaction").(*xorm.Session)
	if tx != nil {
		tx.Close()
	}
}
