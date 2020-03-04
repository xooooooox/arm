package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
	"time"
)

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", Args.DbHost, Args.DbPort, Args.DbUser, Args.DbName, Args.DbPass)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(time.Minute)
	go func(db *sql.DB) {
		for {
			err := db.Ping()
			if err != nil {
				log.Println("database ping error: ", err.Error())
			}
			time.Sleep(time.Second * 5)
		}
	}(db)
	DB = db
}

type Tables struct {
	Name    string
	Comment string
}

func PgTables() (result []Tables) {
	query := "select relname as name,cast(obj_description(relfilenode,'pg_class') as varchar) as comment from pg_class c where relname in (select tablename from pg_tables where schemaname='public' and position('_2' in tablename)=0);"
	rows, err := DB.Query(query)
	if err != nil {
		log.Println("sql error:", err.Error())
	}
	tmp := Tables{}
	for rows.Next() {
		err = rows.Scan(&tmp.Name, &tmp.Comment)
		if err != nil {
			log.Println("sql error:", err.Error())
		}
		result = append(result, tmp)
	}
	return result
}

type Columns struct {
	Num     int64
	Name    string
	Type    string
	Comment string
}

func PgColumns(table string) (result []Columns) {
	rows, err := DB.Query(`select a.attnum as num,a.attname as name,concat_ws('',t.typname,SUBSTRING(format_type(a.atttypid,a.atttypmod) from '\(.*\)')) as type,d.description as comment from pg_class c,pg_attribute a,pg_type t,pg_description d where c.relname=$1 and a.attnum>0 and a.attrelid=c.oid and a.atttypid=t.oid and d.objoid=a.attrelid and d.objsubid=a.attnum;`, table)
	if err != nil {
		log.Println("sql error:", err.Error())
	}
	tmp := Columns{}
	for rows.Next() {
		err = rows.Scan(&tmp.Num, &tmp.Name, &tmp.Type, &tmp.Comment)
		if err != nil {
			log.Println("sql error:", err.Error())
		}
		result = append(result, tmp)
	}
	return result
}

func PgTypeToGoType(p string) (g string) {
	if strings.Index(p, "numeric") >= 0 || strings.Index(p, "decimal") >= 0 {
		return "float64"
	}
	switch p {
	case "bigint", "int8", "bigserial", "serial8":
		g = "int64"
	case "int", "integer", "int4", "serial", "serial4":
		g = "int"
	case "smallint", "int2", "smallserial", "serial2":
		g = "int16"
	case "double", "money", "float8":
		g = "float64"
	case "real":
		g = "float"
	case "bytea", " ":
		g = "[]byte"
	case "boolean":
		g = "bool"
	default:
		g = "string"
	}
	return g
}
