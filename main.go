package main

import (
  "github.com/jubnzv/go-taskwarrior"
)

func main() {
  tw, _ := taskwarrior.NewTaskWarrior("~/.taskrc")
  tw.FetchAllTasks()
  tw.PrintTasks()
}
