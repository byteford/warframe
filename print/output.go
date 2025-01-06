package print

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Printf(format string, a ...any) (int, error) {
	return fmt.Printf(format, a...)
}

func Output(format string, a ...any) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, format, a...)
	fmt.Fprintf(w, "-----\n")
	w.Flush()

}
