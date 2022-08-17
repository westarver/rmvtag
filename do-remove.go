package rmvtag

import (
	"fmt"
	"os"
	"strings"

	"github.com/westarver/helper"
)

//─────────────┤ doRemove ├─────────────

func doRemove(in, out string, rng []int, mock bool) error {
	slice, output, ranges, err := getCmdParameters(in, out, rng, mock)
	if output != os.Stdout {
		defer output.Close()
	}

	if err != nil {
		return err
	}

	var start, end int
	var sb strings.Builder

	start = ranges[0][0]
	if start > 0 {
		for i := 0; i < start; i++ { // since ranges are sorted the lines
			sb.WriteString(slice[i] + "\n") // between 0 and start of first range
		} // are guaranteed not to be in a range
	}

	for i := 0; i < len(ranges); i++ { //possibly 1 range encompassing all
		start = ranges[i][0]
		end = ranges[i][1]

		if end > len(slice) {
			end = len(slice)
		}

		sb.WriteString(removeTaggedLines(slice[start:end]))

		var next = len(slice)
		if end < next {
			if len(ranges) > i+1 {
				next = ranges[i+1][0]
			}

			if next > len(slice) {
				next = len(slice)
			}

			for _, l := range slice[end:next] {
				sb.WriteString(l + "\n")
			}
			end = next
		}
	}
	if end < len(slice) { // output the unaffected after range lines last
		for _, l := range slice[end:] {
			sb.WriteString(l + "\n")
		}
	}
	// write the string builder contents to the output
	lines := helper.GetLinesFromString(sb.String())
	for _, line := range lines {
		if line == "~~[rmv]" {
			if mock {
				fmt.Fprintln(output, line)
			}
		} else {
			fmt.Fprintln(output, line)
		}
	}

	return nil
}
