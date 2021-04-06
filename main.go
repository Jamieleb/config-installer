package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type path string
type dir string

func main() {
  targetDir := dir("/.config/nvim")
  home, hmErr := os.UserHomeDir()
  if hmErr != nil {
    fmt.Println("Error:", hmErr)
    os.Exit(1)
  }

  files := findNvimConfFiles(home, regexp.MustCompile(`.*\.(vim|lua|yaml)$`))
  dirsToCreate := getDirsFromPaths(files)
  removeOldLinksAndDirs(files, dirsToCreate, home, targetDir)
  createDirs(dirsToCreate, home, targetDir)
  createLinks(files, home, targetDir)
  fmt.Println(files)
  fmt.Println(dirsToCreate)
}

func createDirs(dirsToCreate []dir, home string, target dir) {
  for _, d := range dirsToCreate {
    mkdirErr := os.MkdirAll(home + "/" + string(target) + "/" + string(d), 0755)
    if mkdirErr != nil {
      fmt.Println(mkdirErr)
      os.Exit(1)
    }
  }
}

func createLinks(files []path, home string, target dir) []error {
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

func removeOldLinksAndDirs(paths []path, dirs []dir, h string, td dir) ([]error, []error) {
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

func findNvimConfFiles(home string, re *regexp.Regexp) []path {
  ps := []path{}
  os.Chdir(home + "/config")

  err := filepath.Walk(".", func(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    if re.FindString(info.Name()) != "" {
      ps = append(ps, path(p))
    }
    return nil
  })
  if err != nil {
      log.Println(err)
  }
  return ps
}

func strInSlice(slice []string, str string) bool {
  for _, s := range slice {
    if s == str {
      return true
    }
  }
  return false
}

func getDirsFromPaths(paths []path) []dir {
  keys := make(map[string]bool)
  uniqDirs := []dir{}

  for _, fp := range paths {
    if d := filepath.Dir(string(fp)); !keys[d] {
      keys[d] = true
      // We don't want the root directory as it already exists
      if d != "." {
        uniqDirs = append(uniqDirs, dir(d))
      }
    }
  }

  return uniqDirs
}
