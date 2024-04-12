package controller

import (
    "bufio"
    "fmt"
    "gosha/internal/storage"
    "os"
    "strings"
)

func clear() {
    store, err := storage.GetStore()
    if err != nil {
        panic(err)
    }

    list()

    fmt.Println("Are you sure? (y/n)")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

    fmt.Println()
    if formattedInput := strings.Trim(strings.ToLower(input), "\n")[0]; formattedInput != 'y' {
        fmt.Println("It's ok, you can do it later :)")
        os.Exit(0)
    }

    store.Truncate(0)
    fmt.Println("Clear!")
}
