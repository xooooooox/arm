package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func WriteFile(file string, s *string) error {
	ds := string(filepath.Separator)
	abs, _ := filepath.Abs(Args.FileSaveDir)
	file = abs + ds + file
	_, err := os.Stat(file)
	// file exist
	if err == nil {
		// delete this file
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}
	// create file
	f, err := os.Create(file)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return err
	}
	_, err = f.WriteString(*s)
	if err != nil {
		return err
	}
	return FmtGoFile(file)
}

func FmtGoFile(file string) error {
	cmd := exec.Command("go", "fmt", file)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
