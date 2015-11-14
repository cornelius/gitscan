package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "github.com/olekukonko/tablewriter"
  "path"
  "sort"
)

type GitRepo struct {
  name string
  branch string
  status string
  path string
}

func findGitRepositories(dir string, gitRepos map[string]GitRepo) {
  files, _ := ioutil.ReadDir(dir)
  for _, f := range files {
    dirName := dir + "/" + f.Name()
    if f.IsDir() {
      findGitRepositories(dirName, gitRepos)
      if f.Name() == ".git" {
        gitRepos[dir] = GitRepo{name: path.Base(dir), branch: GitBranch(dir), status: GitStatus(dir), path: path.Dir(dir)}
      }
    }
  }
}

func InspectGit(root string) {
  if _, err := os.Stat(root); err == nil {
    fmt.Printf("Inspecting system for git repos at '%v'...\n", root)

    gitRepos := make(map[string]GitRepo)
    findGitRepositories(root, gitRepos)
    if len(gitRepos) >= 0 {
      keys := make([]string, len(gitRepos))
      i := 0
      for key := range gitRepos {
        keys[i] = key
        i++
      }
      sort.Strings(keys)

      table := tablewriter.NewWriter(os.Stdout)
      table.SetHeader([]string{"Name", "Branch", "Status", "Path"})

      for i := range keys {
        repo := gitRepos[keys[i]]
        v := []string{repo.name, repo.branch, repo.status, repo.path}
        table.Append(v)
      }
      table.Render()
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
