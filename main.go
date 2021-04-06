package main

import (
	"fmt"
	"regexp"
	"strings"
)

type pathList []string
type filepaths pathList
type directories pathList
type message string

func main() {
  targetDir := "~/.config/nvim"
  sourceDir := "~/config"

  fmt.Println(files)
  fmt.Println(dirsToCreate)
}

func neovimSetup(src string, target string) error {
  clearMessage := printMessage("Walking 'source' tree to locate Neovim configuration files...")
  files := findNvimConfFiles(src, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  clearMessage()

  clearMessage = printMessage("Creating required directories in " + target)
  dirsToCreate := files.extractDirs()
  clearMessage()

  pathList(files).removeFilesOrDirs(target)
  pathList(dirsToCreate).removeFilesOrDirs(target)

  dirsToCreate.create(target)

  symlinkErrors := files.createSymLinks(src, target)
  if len(symlinkErrors) > 0 {
    for _, e := range symlinkErrors { fmt.Println("Error:", e) }
  }


}

func printMessage(msg string) func() {
  fmt.Print(msg)
  return func() { fmt.Print(strings.Repeat("\b", len(msg))) }
}
