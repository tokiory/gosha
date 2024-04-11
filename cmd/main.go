package main

import (
	"flag"
	"fmt"
	"gosha/internal/arguments"
	"gosha/internal/controller"
	"os"
)

func main() {
  userArgs := append(os.Args, "", "")
  userArgs = userArgs[2:]

  flag.Parse()

  switch {
    case arguments.IsHelpFlag(): 
    controller.Run(controller.CommandHelp)

    case arguments.IsAddFlag(): 
      if userArgs[0] == "" {
        fmt.Println(arguments.MissingArgMessage("name"))
        os.Exit(1)
      }

      controller.Run(controller.CommandAdd, userArgs[0])
    
    case arguments.IsRemoveFlag():
      if userArgs[0] == "" {
        fmt.Println(arguments.MissingArgMessage("id"))
        os.Exit(1)
      }

      fmt.Println("Remove an item!")

    case arguments.IsModFlag():
      if userArgs[0] == "" {
        fmt.Println(arguments.MissingArgMessage("id"))
        os.Exit(1)
      }

      if userArgs[1] == "" {
        fmt.Println(arguments.MissingArgMessage("name"))
        os.Exit(1)
      }

      fmt.Println("Change an item!")

    case arguments.IsCheckFlag():
      if userArgs[0] == "" {
        fmt.Println(arguments.MissingArgMessage("id"))
        os.Exit(1)
      }

      controller.Run(controller.CommandCheck, userArgs[0])

  default:
    controller.Run(controller.CommandList)
  }
}
