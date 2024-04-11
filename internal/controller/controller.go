package controller

func Run(command int, args ...string) {
  switch command {
  case CommandAdd:
    add(args[0])
  case CommandHelp:
    help()
  case CommandCheck:
    check(args[0])
  case CommandModify:
    // TODO: Modify logic
  case CommandRemove:
    // TODO: Remove logic
  case CommandList:
    list()
  }
}
