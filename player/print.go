package player

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func Print(p Player) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "-----\n")
	fmt.Fprintf(w, "%+v\n", p)
	fmt.Fprintf(w, "-----\n")
	w.Flush()
}
