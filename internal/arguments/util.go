package arguments

import (
	"os"
)

func checkArgument(arg string) bool {
  return len(os.Args) >= 2 && os.Args[1] == arg
}
