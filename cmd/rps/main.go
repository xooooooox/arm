package main

import (
	"database/sql"
	"fmt"
	uf "github.com/xooooooox/arm/utils/file"
	un "github.com/xooooooox/arm/utils/name"
	"log"
	"strings"
)

var (
	Args CmdArgs
	DB   *sql.DB
)

func main() {
	if err := Write(); err != nil {
		log.Fatalln(err.Error())
	}
}

func Write() error {
	s := ""
	for _, t := range PgTables() {
		TableName := un.UnderlineToPascal(t.Name)
		s = fmt.Sprintf("%s\n// %s %s \n type %s struct {\n", s, TableName, t.Comment, TableName)
		for _, c := range PgColumns(t.Name) {
			ColumnName := un.UnderlineToPascal(c.Name)
			s = fmt.Sprintf("%s\t%s %s ", s, ColumnName, PgTypeToGoType(c.Type))
			if Args.JsonTag {
				if !strings.HasSuffix(s, "`") {
					s = fmt.Sprintf("%s`", s)
				}
				s = fmt.Sprintf("%s%s`", s, JsonTag(c))
			}
			s = fmt.Sprintf("%s // %s\n", s, c.Comment)
		}
		s = fmt.Sprintf("%s}\n", s)
	}
	fileTmp := `package %s
%s
`
	s = fmt.Sprintf(fileTmp, Args.FilePkgName, s)
	filename := fmt.Sprintf("%s%s%s.go", Args.FileSaveDir, Args.DbName, Args.FileSuffixName)
	_, err := uf.WriteToFile(&s, filename, Args.FileSaveDir)
	if err != nil {
		return err
	}
	return uf.Fmt(Args.FileSaveDir + filename)
}

// JsonTag
func JsonTag(c Columns) string {
	ignores := strings.Split(Args.JsonIgnore, ",")
	column := un.PascalToUnderline(c.Name)
	for _, v := range ignores {
		if column == v {
			column = "-"
		}
	}
	return fmt.Sprintf("json:\"%s\"", column)
}
