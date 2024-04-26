package database

import (
	"brew-note/src/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CloseDBConnectionFunc is func type
type CloseDBConnectionFunc func()

// SQLHandler is struct
type SQLHandler struct {
	Conn  *gorm.DB
	Close CloseDBConnectionFunc
}

var Handler SQLHandler

// NewSQLHandler is connect to database
func NewSQLHandler() {
	fmt.Print("NewSQLHandler")
	db := gormConnect()
	DB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	DB.SetMaxIdleConns(config.DB.MaxIdleConns)
	DB.SetMaxOpenConns(config.DB.MaxOpenConns)
	tempHandler, err := &SQLHandler{
		Conn: db,
		Close: func() {
			if db != nil {
				if err := DB.Close(); err != nil {
					panic(err.Error())
				}
			}
		},
	}, err
	if err != nil {
		panic(err.Error())
	}
	Handler = *tempHandler
}

func gormConnect() *gorm.DB {
	fmt.Print("gormConnect")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	fmt.Print(connStr)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

// Find find records that match given conditions
func (h *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.Find(out, where...)
}

// First find firs record that match given conditions, order by primary key
func (h *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.First(out, where...)
}

// Create insert the value into database
func (h *SQLHandler) Create(value interface{}) *gorm.DB {
	return h.Conn.Create(value)
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (h *SQLHandler) Save(value interface{}) *gorm.DB {
	return h.Conn.Save(value)
}

// Updates update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
func (h *SQLHandler) Updates(values interface{}) *gorm.DB {
	return h.Conn.Updates(values)
}

// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition WARNING If model has DeletedAt field, GORM will only set field DeletedAt's value to current time
func (h *SQLHandler) Delete(value interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.Delete(value, where...)
}

// Where return a new relation, filter records with given conditions, accepts `map`, `struct` or `string` as conditions, refer http://jinzhu.github.io/gorm/crud.html#query
func (h *SQLHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return h.Conn.Where(query, args...)
}

// FirstOrCreate find first matched record or create a new one with given conditions (only works with struct, map conditions) https://jinzhu.github.io/gorm/crud.html#firstorcreate
func (h *SQLHandler) FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB {
	return h.Conn.FirstOrCreate(out, where...)
}

// Select specify fields that you want to retrieve from database when querying, by default, will select all fields; When creating/updating, specify fields that you want to save to database
func (h *SQLHandler) Select(query interface{}, args ...interface{}) *gorm.DB {
	return h.Conn.Select(query, args...)
}
