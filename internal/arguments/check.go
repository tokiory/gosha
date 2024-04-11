package arguments

import (
	"flag"
	"strings"
)

var CheckFlagShort = flag.Bool(strings.Split(checkKeyword, "")[0], false, "")
var CheckFlag = flag.Bool(checkKeyword, false, "")

func IsCheckFlag() bool {
  command := checkArgument(checkKeyword)
  return command || *CheckFlag || *CheckFlagShort
}
