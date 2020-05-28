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
	query := `SELECT "relname" AS "name", COALESCE(cast(obj_description("relfilenode",'pg_class') AS VARCHAR),'') AS "comment" FROM "pg_class" "c" WHERE ( "relname" IN ( SELECT "tablename" FROM "pg_tables" WHERE ( "schemaname" = $1 AND position('_2' IN "tablename") = 0 ) ) ) ORDER BY "name" ASC;`
	rows, err := DB.Query(query, Args.DbSchemaName)
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
	rows, err := DB.Query(`SELECT "a"."attnum" AS "num", "a"."attname" AS "name", concat_ws('', "t"."typname", SUBSTRING(format_type("a"."atttypid", "a"."atttypmod") FROM '\(.*\)')) AS "type", "d"."description" AS "comment" FROM "pg_class" "c", "pg_attribute" "a", "pg_type" "t", "pg_description" "d" WHERE ( "c"."relname" = $1 AND "a"."attnum" > 0 AND "a"."attrelid"="c"."oid" AND "a"."atttypid" = "t"."oid" AND "d"."objoid" = "a"."attrelid" AND "d"."objsubid" = "a"."attnum" ) ORDER BY "num" ASC;`, table)
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
