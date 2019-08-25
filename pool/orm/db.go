package orm

import (
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"


type Orm interface {
	InitDB()
	Find()
	FindAll()
	Insert()
	Update()
	Delete()
}

type DB struct {
	Orm *gorm.DB
}

func (db *DB) InitDB(dsn string)  (err error) {
	db.Orm,err= gorm.Open("mysql",dsn)
	if err!=nil{

		//log.Fatal(err)
		return
	}
	db.Orm.SingularTable(true)
	return nil
}

func (db *DB) Find(sql string) (result interface{}) {

	db.Orm.Raw(sql).Scan(&result)
	return
}