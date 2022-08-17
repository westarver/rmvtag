package rmvtag

import (
	"regexp"
	"strings"
)

//─────────────┤ removeTaggedLines ├─────────────

func removeTaggedLines(slice []string) string {
	linepat := `[ \t]*[.]*<rmv/>`
	linetag := regexp.MustCompile(linepat)
	openpat := `[ \t]*<rmv>`
	opentag := regexp.MustCompile(openpat)
	closepat := `[ \t]*</rmv>`
	closetag := regexp.MustCompile(closepat)

	inblk := false
	var sb = strings.Builder{}
	for _, line := range slice {
		if linetag.MatchString(line) {
			inblk = false
			sb.WriteString("~~[rmv]\n")
			continue
		}
		if opentag.MatchString(line) {
			inblk = true
			sb.WriteString("~~[rmv]\n")
			continue
		}
		if closetag.MatchString(line) {
			inblk = false
			sb.WriteString("~~[rmv]\n")
			continue
		}
		if inblk {
			sb.WriteString("~~[rmv]\n")
		} else {
			sb.WriteString(line + "\n")
		}
	}

	return sb.String()
}
