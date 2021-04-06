package main

import (
	"fmt"
	"os"
)

func (ps pathList) removeFilesOrDirs(targetDir string) []error {
  errors:= []error{}

  for _, p := range ps {
    err := os.Remove(targetDir + "/" + p)
    if err != nil { errors = append(errors, err) }
  }

  return errors
}

func (dirs directories) create(targetDir string) {
  for _, d := range dirs {
    err := os.MkdirAll(targetDir + "/" + d, 0755)
    if err != nil {
      fmt.Println("Error:", err)
      os.Exit(1)
    }
  }
}

func (files filepaths) createSymLinks(src string, target string) []error {
  errors := []error{}

  for _, f := range files {
    err := os.Symlink(getFullPath(src, f), getFullPath(target, f))
    if err != nil { errors = append(errors, err) }
  }

  return errors
}

func getFullPath(base string, filepath string) string {
  return base + "/" + filepath
}
