package arguments

import (
	"flag"
	"strings"
)

var AddFlagShort = flag.Bool(strings.Split(addKeyword, "")[0], false, "")
var AddFlag = flag.Bool(addKeyword, false, "")

func IsAddFlag() bool {
  command := checkArgument(addKeyword)
  return command || *AddFlag || *AddFlagShort
}
