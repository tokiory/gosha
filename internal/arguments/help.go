package arguments

import (
	"flag"
  "strings"
)

const HelpMessage = `Gosha Help:
help, --help, -h:                Show current message
add <name>, -a <name>:           Add new todo item
remove <id>, -r <id>:            Remove todo item
mod <id> <name>, -m <id> <name>: Modify todo item
check <id>, -c <id>:             Toggle todo item
<nothing>:                       List not checked todo items
all:                             List all todo items`

var HelpFlagShort = flag.Bool(strings.Split(helpKeyword, "")[0], false, "")
var HelpFlag = flag.Bool(helpKeyword, false, "")

func IsHelpFlag() bool {
  command := checkArgument(helpKeyword)
  return command || *HelpFlag || *HelpFlagShort
}
