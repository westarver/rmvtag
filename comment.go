package rmvtag

import (
	"regexp"
	"strings"

	"github.com/westarver/helper"
)

//─────────────┤ commentTaggedLines ├─────────────

func commentTaggedLines(slice []string) string {
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
			ws := helper.LeadingWs(line)
			line = helper.StripComment(line, CommentString)
			line = CommentString + line
			sb.WriteString(ws + line + "\n")
			inblk = false
			continue
		}
		if opentag.MatchString(line) {
			ws := helper.LeadingWs(line)
			line = helper.StripComment(line, CommentString)
			line = CommentString + line
			sb.WriteString(ws + line + "\n")
			inblk = true
			continue
		}
		if closetag.MatchString(line) {
			ws := helper.LeadingWs(line)
			line = helper.StripComment(line, CommentString)
			line = CommentString + line
			sb.WriteString(ws + line + "\n")
			inblk = false
			continue
		}
		if inblk {
			ws := helper.LeadingWs(line)
			line = helper.StripComment(line, CommentString)
			line = CommentString + line
			sb.WriteString(ws + line + "\n")
		} else {
			sb.WriteString(line + "\n")
		}
	}
	return sb.String()
}
