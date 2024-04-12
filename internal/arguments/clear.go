package arguments

func IsClearFlag() bool {
    command := checkArgument(clearKeyword)
    return command
}
