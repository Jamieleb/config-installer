package main

import "os"

func (ps filepaths) removeFilesOrDirs(targetDir string) []error {
  errors:= []error{}

  for _, p := range ps {
    err := os.Remove(targetDir + "/" + p)
    if err != nil { errors = append(errors, err) }
  }

  return errors
}
