package pk

import (
	"database/sql"
	"fmt"
	"log"
	"polaroid/config"
)

type Pk struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func NewPk(s config.Conf) (a *Pk) {
	a = &Pk{
		Host:     s["host"].(string),
		Port:     s["port"].(int),
		User:     s["user"].(string),
		Password: s["password"].(string),
		Dbname:   s["dbname"].(string),
	}
	a.Starter()
	return

}

func (a *Pk) Starter() {
	db, err := a.Connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS article (
		title VARCHAR(50) UNIQUE,
		describe VARCHAR(50), 
		logo VARCHAR(50),
		body VARCHAR(50), 
		css VARCHAR(50),
		name VARCHAR(50
	) UNIQUE);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS gallerie (
		title VARCHAR(50) UNIQUE,
		describe VARCHAR(50),
		logo VARCHAR(50),
		body VARCHAR(50),
		css VARCHAR(50),
		name VARCHAR(50) UNIQUE
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *Pk) Connect() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", a.Host, a.Port, a.User, a.Password, a.Dbname)
	if db, err = sql.Open("postgres", psqlInfo); err != nil {
		return
	}
	return db, db.Ping()
}
