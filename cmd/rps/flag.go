package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CmdArgs struct {
	DbHost         string
	DbPort         int
	DbUser         string
	DbPass         string
	DbName         string
	DbSchemaName   string
	Version        bool
	FilePkgName    string
	FileSaveDir    string
	FileSuffixName string
	JsonTag        bool
}

var Version string

func init() {
	Version = "0.1.0"
	osArgs := os.Args
	for i := 0; i < len(osArgs); i++ {
		if osArgs[i] == "-v" || osArgs[i] == "--version" {
			fmt.Println(Version)
			os.Exit(0)
		}
	}
	flag.BoolVar(&Args.Version, "v", false, "view version")
	flag.StringVar(&Args.DbHost, "h", "127.0.0.1", "database host name")
	flag.IntVar(&Args.DbPort, "P", 5432, "database host port")
	flag.StringVar(&Args.DbUser, "u", "postgres", "database connect user")
	flag.StringVar(&Args.DbPass, "p", "postgres", "database connect user password")
	flag.StringVar(&Args.DbName, "n", "postgres", "database name")
	flag.StringVar(&Args.DbSchemaName, "s", "public", "database schema name")
	flag.StringVar(&Args.FilePkgName, "N", "dao", "package name")
	flag.StringVar(&Args.FileSaveDir, "d", "./", "output file disk address")
	flag.StringVar(&Args.FileSuffixName, "S", "_tmp", "file suffix name")
	flag.BoolVar(&Args.JsonTag, "j", false, "json tag")
	flag.Parse()
	// make sure dir name use filepath.Separator end
	ds := string(filepath.Separator)
	if !strings.HasSuffix(Args.FileSaveDir, ds) {
		Args.FileSaveDir = fmt.Sprintf("%s%s", Args.FileSaveDir, ds)
	}
}
