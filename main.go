package main

import (
	"fmt"
	"regexp"
)

type pathList []string
type filepaths pathList
type directories pathList
type message string

func main() {
  targetDir := "~/.config/nvim"
  sourceDir := "~/config"

  neovimSetup(sourceDir, targetDir)
}

func neovimSetup(src string, target string) {
  clearMessage := printMessage("Walking 'source' tree to locate Neovim configuration files...")
  files := findNvimConfFiles(src, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  clearMessage()

  dirsToCreate := files.extractDirs()

  clearMessage = printMessage("Removing old files and directories...")
  pathList(files).removeFilesOrDirs(target)
  pathList(dirsToCreate).removeFilesOrDirs(target)
  clearMessage()

  clearMessage = printMessage("Creating required directories in " + target)
  dirsToCreate.create(target)
  clearMessage()

  clearMessage = printMessage("Creating SymLinks for Neovim configuration files...")
  symlinkErrors := files.createSymLinks(src, target)
  if len(symlinkErrors) > 0 {
    for _, e := range symlinkErrors { fmt.Println("Error:", e) }
  } else {
    clearMessage()
  }
}
