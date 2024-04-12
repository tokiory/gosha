package controller

func Run(command int, args ...string) {
    switch command {
    case CommandAll:
        all()
    case CommandClear:
        clear()
    case CommandAdd:
        add(args[0])
    case CommandHelp:
        help()
    case CommandCheck:
        check(args[0])
    case CommandModify:
        modify(args[0], args[1])
    case CommandRemove:
        remove(args[0])
    case CommandList:
        list()
    }
}
