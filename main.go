package main

import (
	"fmt"
	"regexp"
)

type pathList []string
type filepaths pathList
type directories pathList

func main() {
  targetDir := "~/.config/nvim"
  sourceDir := "~/config"

  files := findNvimConfFiles(sourceDir, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  dirsToCreate := files.extractDirs()

  pathList(files).removeFilesOrDirs(targetDir)
  pathList(dirsToCreate).removeFilesOrDirs(targetDir)

  dirsToCreate.create(targetDir)

  symlinkErrors := files.createSymLinks(sourceDir, targetDir)
  if len(symlinkErrors) > 0 {
    for _, e := range symlinkErrors { fmt.Println("Error:", e) }
  }

  fmt.Println(files)
  fmt.Println(dirsToCreate)
}
