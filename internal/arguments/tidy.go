package arguments

func IsTidyFlag() bool {
	command := checkArgument(tidyKeyword)
	return command
}
