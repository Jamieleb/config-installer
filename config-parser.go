package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type filepaths []string

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

func (fps filepaths) extractDirs() filepaths {
  encounteredDirs := make(map[string]bool)
  uniqDirs := filepaths{}

  for _, p := range fps {
    if d := filepath.Dir(p); !encounteredDirs[d] {
      encounteredDirs[d] = true
      if d != "." {
        uniqDirs = append(uniqDirs, d)
      }
    }
  }

  return uniqDirs
}
