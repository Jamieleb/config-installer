package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
  targetDir := "/.config/nvim"
  home, hmErr := os.UserHomeDir()
  if hmErr != nil {
    fmt.Println("Error:", hmErr)
    os.Exit(1)
  }

  files := findNvimConfFiles(home, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  dirsToCreate := files.extractDirs()
  removeOldLinksAndDirs(files, dirsToCreate, home, targetDir)
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

func removeOldLinksAndDirs(paths filepaths, dirs filepaths, h string, td string) ([]error, []error) {
  pathErrs := []error{}
  dirErrs := []error{}

  for _, p := range paths {
    err := os.Remove(h + string(td) + "/" + string(p))
    if err != nil {
      pathErrs = append(pathErrs, err)
    }
  }

  for _, d := range dirs {
    err := os.Remove(h + string(td) + "/" + string(d))
    if err != nil {
      dirErrs = append(dirErrs, err)
    }
  }

  return pathErrs, dirErrs
}
