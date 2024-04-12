package table

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Writer() *tabwriter.Writer {
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	return writer
}

func Concat(writer *tabwriter.Writer, values string) {
	fmt.Println(writer, values)
}

func Flush(writer *tabwriter.Writer) {
	writer.Flush()
}
