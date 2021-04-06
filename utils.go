package main

import (
	"fmt"
	"strings"
)

func printMessage(msg string) func() {
  fmt.Print(msg)
  return func() { fmt.Print(strings.Repeat("\b", len(msg))) }
}
