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
    fmt.Fprintf(os.Stderr, "%v: Error creating output pipe for '%v': %v\n", repo, cmdString, err)
    os.Exit(1)
  }

  branch := "unknown branch"

  scanner := bufio.NewScanner(out)

  err = cmd.Start()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v: Error starting '%v': %v\n", repo, cmdString, err)
    os.Exit(1)
  }

  for scanner.Scan() {
    branch = scanner.Text()
  }

  err = cmd.Wait()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v: Error waiting for '%v': %v\n", repo, cmdString, err)
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
    fmt.Fprintf(os.Stderr, "%v: Error creating output pipe for '%v': %v\n", repo, cmdString, err)
    os.Exit(1)
  }

  status := "unknown status"

  scanner := bufio.NewScanner(out)

  err = cmd.Start()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v: Error starting '%v': %v\n", repo, cmdString, err)
    os.Exit(1)
  }

  for scanner.Scan() {
    line := scanner.Text()
    if !strings.HasPrefix(line, "#") {
      status = line
    }
  }

  err = cmd.Wait()
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v: Error waiting for '%v': %v\n", repo, cmdString, err)
    os.Exit(1)
  }

  return status
}
