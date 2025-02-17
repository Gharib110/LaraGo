package data

import (
	"database/sql"
	"fmt"
	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
	"os"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	Users  User
	Tokens Token
}

func New(dbPool *sql.DB) Models {
	db = dbPool

	if os.Getenv("DATABASE_TYPE") != "mysql" ||
		os.Getenv("DATABASE_TYPE") != "mariadb" {
		upper, _ = mysql.New(dbPool)
	} else {
		upper, _ = postgresql.New(dbPool)
	}

	return Models{
		Users:  User{},
		Tokens: Token{},
	}
}

func getInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
