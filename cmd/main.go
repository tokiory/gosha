package main

import (
	"flag"
	"gosha/internal/arguments"
	"gosha/internal/controller"
	"log"
	"os"
)

func main() {
	userArgs := append(os.Args, "", "")
	userArgs = userArgs[2:]

	flag.Parse()

	switch {
	case arguments.IsHelpFlag():
		controller.Run(controller.CommandHelp)

	case arguments.IsClearFlag():
		controller.Run(controller.CommandClear)

	case arguments.IsAddFlag():
		if userArgs[0] == "" {
			log.Fatal(arguments.MissingArgMessage("name"))
		}

		controller.Run(controller.CommandAdd, userArgs[0])

	case arguments.IsRemoveFlag():
		if userArgs[0] == "" {
			log.Fatal(arguments.MissingArgMessage("id"))
		}

		controller.Run(controller.CommandRemove, userArgs[0])

	case arguments.IsModFlag():
		if userArgs[0] == "" {
			log.Fatal(arguments.MissingArgMessage("id"))
		}

		if userArgs[1] == "" {
			log.Fatal(arguments.MissingArgMessage("name"))
		}

		controller.Run(controller.CommandModify, userArgs[0], userArgs[1])

	case arguments.IsCheckFlag():
		if userArgs[0] == "" {
			log.Fatal(arguments.MissingArgMessage("id"))
		}

		controller.Run(controller.CommandCheck, userArgs[0])

	case arguments.IsAllFlag():
		controller.Run(controller.CommandAll)

	case arguments.IsTidyFlag():
		controller.Run(controller.CommandTidy)

	default:
		if len(os.Args) >= 2 && os.Args[1] != "" {
			log.Fatalf("No such command %s", os.Args[1])
		}

		controller.Run(controller.CommandList)
	}
}
