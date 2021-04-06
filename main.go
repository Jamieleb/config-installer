package main

import (
	"fmt"
	"os"
	"regexp"
)

type pathList []string
type filepaths pathList
type directories pathList

func main() {
  targetDir := "/.config/nvim"
  sourceDir := "~/config"

  files := findNvimConfFiles(sourceDir, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  dirsToCreate := files.extractDirs()

  pathList(files).removeFilesOrDirs(targetDir)
  pathList(dirsToCreate).removeFilesOrDirs(targetDir)

  createDirs(dirsToCreate, home, targetDir)
  createLinks(files, home, targetDir)
  fmt.Println(files)
  fmt.Println(dirsToCreate)
}

func createDirs(dirsToCreate filepaths, home string, target string) {
  for _, d := range dirsToCreate {
    mkdirErr := os.MkdirAll(home + "/" + string(target) + "/" + string(d), 0755)
    if mkdirErr != nil {
      fmt.Println(mkdirErr)
      os.Exit(1)
    }
  }
}

func createLinks(files filepaths, home string, target string) []error {
  errors := []error{}
  for _, f := range files {
    err := os.Symlink(home + "/config/" + string(f), home + string(target) + "/" + string(f))
    if err != nil {
      fmt.Println("Error: ", err)
      errors = append(errors, err)
    }
  }
  return errors
}

