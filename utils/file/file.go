package file

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// WriteToFile Delete the file if it exists
func WriteToFile(content *string, filename string, dirname string) (n int, err error) {
	ds := string(filepath.Separator)
	abs, err := filepath.Abs(dirname)
	if err != nil {
		return
	}
	if !strings.HasSuffix(dirname, ds) {
		abs = fmt.Sprintf("%s%s", abs, ds)
	}
	filename = fmt.Sprintf("%s%s", abs, filename)
	_, err = os.Stat(filename)
	if err == nil {
		err = os.Remove(filename)
		if err != nil {
			return
		}
	}
	f, err := os.Create(filename)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return
	}
	return f.WriteString(*content)
}

// Fmt Fmt go source file
func Fmt(file string) error {
	cmd := exec.Command("go", "fmt", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
