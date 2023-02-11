package main
import (
  "fmt"
  "os"
  "os/user"
  "treewalker/repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hello %s! This is the TreeWalker programming language.\n",user.Username)
  fmt.Printf("Please enter your commands\n")
  repl.Start(os.Stdin, os.Stdout)
}
