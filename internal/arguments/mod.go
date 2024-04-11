package arguments

import (
	"flag"
	"strings"
)

var ModFlagShort = flag.Bool(strings.Split(modifyKeyword, "")[0], false, "")
var ModFlag = flag.Bool(modifyKeyword, false, "")

func IsModFlag() bool {
  command := checkArgument(modifyKeyword)
  return command || *ModFlag || *ModFlagShort
}
