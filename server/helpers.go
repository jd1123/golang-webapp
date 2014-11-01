package server

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func GetCWD() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func ProjectPath(p string) string {
	homedir = GetCWD()
	return path.Join(homedir, p)
}

func TemplatePath(templateName string) string {
	return path.Join(ProjectPath(CONFIG["templateDir"]), templateName)
}

func StaticPath(filename string) string {
	return path.Join(ProjectPath(CONFIG["staticDir"]), filename)
}

func PackFiles(o []os.FileInfo) []string {
	s := make([]string, 0)
	for i := range o {
		s = append(s, o[i].Name())
	}
	return s
}
