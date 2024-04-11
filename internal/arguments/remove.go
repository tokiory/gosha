package arguments

import (
	"flag"
	"strings"
)

var RemoveFlagShort = flag.Bool(strings.Split(removeKeyword, "")[0], false, "")
var RemoveFlag = flag.Bool(removeKeyword, false, "")

func IsRemoveFlag() bool {
  command := checkArgument(removeKeyword)
  return command || *RemoveFlag || *RemoveFlagShort
}
