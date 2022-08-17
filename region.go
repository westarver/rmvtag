package rmvtag

import (
	"strings"

	"github.com/bitfield/script"
	"github.com/westarver/helper"
)

//─────────────┤ getRangeFromRegion ├─────────────

func getRangeFromRegion(source string, rgns ...string) []int {
	var ret []int
	if len(rgns) == 0 {
		return ret
	}
	source, err := helper.ValidatePath(source)
	if err != nil {
		return nil
	}

	input, err := script.File(source).Slice()
	if err != nil {
		return nil
	}
	lineno := 0
	for _, rgn := range rgns {
		for i := lineno; i < len(input); i++ {
			if strings.Contains(input[i], "<rgn "+rgn+">") {
				ret = append(ret, i)
			} else if strings.Contains(input[i], "</rgn "+rgn+">") {
				ret = append(ret, i)
				lineno = i + 1
			}
		}
	}
	return ret
}
