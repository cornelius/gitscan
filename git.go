package main

import (
  "fmt"
  "os"
  "os/exec"
  "bufio"
  "strings"
)

func GitBranch(repo string) string {
  cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
  cmd.Dir = repo

  cmdString := strings.Join(cmd.Args, " ")

  out, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error creating output pipe for '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  branch := "unknown branch"

  scanner := bufio.NewScanner(out)
  go func() {
    for scanner.Scan() {
      branch = scanner.Text()
    }
  }()

  err = cmd.Start()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error starting '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  err = cmd.Wait()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error waiting for '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  return branch
}

func GitStatus(repo string) string {
  cmd := exec.Command("git", "status")
  cmd.Dir = repo

  cmdString := strings.Join(cmd.Args, " ")

  out, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error creating output pipe for '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  status := "unknown status"

  scanner := bufio.NewScanner(out)
  go func() {
    for scanner.Scan() {
      line := scanner.Text()
      if !strings.HasPrefix(line, "#") {
        status = line
      }
    }
  }()

  err = cmd.Start()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error starting '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  err = cmd.Wait()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error waiting for '%v': %v\n", cmdString, err)
    os.Exit(1)
  }

  return status
}
