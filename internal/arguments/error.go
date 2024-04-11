package arguments

import "fmt"

// Getter for missing argument error message
func MissingArgMessage(arg string) string {
  return fmt.Sprintf("Missing required argument: <%s>", arg)
}
