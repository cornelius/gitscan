package main

import (
  "os"
  "fmt"
)

func main() {
  if len(os.Args) != 2 {
    fmt.Println("Usage: gitscan <directory>")
    os.Exit(1)
  }

  InspectGit(os.Args[1])
}
