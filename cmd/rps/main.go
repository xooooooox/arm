package main

import (
	"database/sql"
	"fmt"
	uf "github.com/xooooooox/arm/utils/file"
	un "github.com/xooooooox/arm/utils/name"
	ut "github.com/xooooooox/arm/utils/time"
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
	consts := "\n"
	for _, t := range PgTables() {
		TableName := un.UnderlineToPascal(t.Name)
		s = fmt.Sprintf("%s\n// %s %s \n type %s struct {\n", s, TableName, t.Comment, TableName)
		for _, c := range PgColumns(t.Name) {
			ColumnName := un.UnderlineToPascal(c.Name)
			consts = fmt.Sprintf("%s\n\t%s = \"%s\" // %s\n",consts,fmt.Sprintf("%s%s",TableName,ColumnName),c.Name,c.Comment)
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

// datetime %s

const (
%s
)
%s
`
	s = fmt.Sprintf(fileTmp, Args.FilePkgName, ut.DateTime(),consts, s)
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
