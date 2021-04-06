package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func findNvimConfFiles(source string, re *regexp.Regexp) filepaths {
  paths := filepaths{}
  os.Chdir(source)

  err := filepath.Walk(source, func(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    if re.FindString(info.Name()) != "" {
      paths = append(paths, p)
    }
    return nil
  })
  if err != nil {
      log.Println(err)
  }
  return paths
}

func (files filepaths) extractDirs() directories {
  encounteredDirs := make(map[string]bool)
  uniqDirs := directories{}

  for _, p := range files {
    if d := filepath.Dir(p); !encounteredDirs[d] {
      encounteredDirs[d] = true
      if d != "." {
        uniqDirs = append(uniqDirs, d)
      }
    }
  }

  return uniqDirs
}
