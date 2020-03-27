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
	for _, t := range PgTables() {
		s = fmt.Sprintf("%s\n// %s %s \n type %s struct {\n", s, un.UnderlineToPascal(t.Name), t.Comment, un.UnderlineToPascal(t.Name))
		for _, c := range PgColumns(t.Name) {
			s = fmt.Sprintf("%s\t%s %s ", s, un.UnderlineToPascal(c.Name), PgTypeToGoType(c.Type))
			if Args.JsonTag {
				if !strings.HasSuffix(s, "`") {
					s = fmt.Sprintf("%s`", s)
				}
				s = fmt.Sprintf("%s %s`", s, JsonTag(c))
			}
			s = fmt.Sprintf("%s // %s\n", s, c.Comment)
		}
		s = fmt.Sprintf("%s}\n", s)
	}
	fileTmp := `package %s

// datetime %s
%s
`
	s = fmt.Sprintf(fileTmp, Args.FilePkgName, ut.DateTime(), s)
	filename := fmt.Sprintf("%s%s%s.go", Args.FileSaveDir, Args.DbName, Args.FileSuffixName)
	_, err := uf.WriteToFile(&s, filename, Args.FileSaveDir)
	if err != nil {
		return err
	}
	return uf.Fmt(Args.FileSaveDir + filename)
}

func JsonTag(c Columns) string {
	return fmt.Sprintf("json:\"%s\"", un.PascalToUnderline(c.Name))
}
