package rmvtag

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//─────────────┤ doList ├─────────────

func doList(in string) error {
	inp, err := os.ReadFile(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "listing tagged lines for", in)
	fmt.Fprintln(os.Stdout, strings.Repeat("-", 80))
	inblk := false
	slice := strings.Split(string(inp), "\n")
	for i, line := range slice {
		line = strings.TrimLeft(line, " \t")
		textLen := len("tagged line #") + len(strconv.Itoa(i))
		colonPad := ":" + strings.Repeat(" ", 24-textLen)

		if strings.Contains(line, LineTag) {
			fmt.Fprintln(os.Stdout, "tagged line #", i, colonPad, line)
			inblk = false
			continue
		}

		if strings.Contains(line, OpenTag) {
			fmt.Fprintln(os.Stdout, "tagged line #", i, colonPad, line)
			inblk = true
			continue
		}

		if strings.Contains(line, CloseTag) {
			fmt.Fprintln(os.Stdout, "tagged line #", i, colonPad, line)
			inblk = false
			continue
		}

		if inblk {
			fmt.Fprintln(os.Stdout, "tag affected line #", colonPad, ": ", line)
		}
	}
	fmt.Fprintln(os.Stdout, strings.Repeat("-", 80))
	return nil
}
