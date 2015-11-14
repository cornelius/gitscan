package main

import (
  "fmt"
  "os"
  "io/ioutil"
)

func findGitRepositories(dir string, gitRepos map[string]string) {
  files, _ := ioutil.ReadDir(dir)
  for _, f := range files {
    dirName := dir + "/" + f.Name()
    if f.IsDir() {
      findGitRepositories(dirName, gitRepos)
      if f.Name() == ".git" {
        gitRepos[dir] = GitStatus(dir) + " (" + GitBranch(dir) + ")"
      }
    }
  }
}

func InspectGit(root string) {
  if _, err := os.Stat(root); err == nil {
    fmt.Printf("Inspecting system for git repos at '%v'...\n", root)

    gitRepos := make(map[string]string)
    findGitRepositories(root, gitRepos)
    if len(gitRepos) >= 0 {
      fmt.Printf("Found git repos:\n")
      for repo := range gitRepos {
        fmt.Printf("  %v: %v\n", repo, gitRepos[repo])
      }
    } else {
      fmt.Printf("No repos found")
    }
  } else {
    if os.IsNotExist(err) {
      fmt.Printf("Directory '%v' does not exist\n", root)
      os.Exit(1)
    }
  }
}

func Inspect(args []string) {
  scope := args[0]
  if scope == "git" {
    var root string
    if len(args) == 2 {
      root = args[1]
    } else {
      root = "/"
    }
    InspectGit(root)
  } else {
    fmt.Printf("Unknown scope '%v'\n", scope)
  }
}
