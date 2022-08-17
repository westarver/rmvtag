package rmvtag

import (
	"os"

	"github.com/westarver/helper"
	ranger "github.com/westarver/ranger"
)

//─────────────┤ getCmdParameters ├─────────────

func getCmdParameters(in, out string, rng []int, mock bool) ([]string, *os.File, [][2]int, error) {
	in, err := helper.ValidatePath(in)
	if err != nil {
		return nil, nil, nil, err
	}
	input, err := os.ReadFile(in)
	if err != nil {
		return nil, nil, nil, err
	}

	inp := string(input)
	var rmv *os.File

	// if mocking then the input stays unchanged and
	// the mock changes go to stdout
	if !mock {
		out, err = helper.ValidatePath(out)
		if err != nil {
			return nil, nil, nil, err
		}

		err = os.WriteFile(out, input, 0666) //save original
		if err != nil {
			return nil, nil, nil, err
		}

		rmv, err = os.OpenFile(in, os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return nil, nil, nil, err
		}
	} else {
		rmv = os.Stdout
	}

	bigSlice := helper.GetLinesFromString(inp)
	var rang []int
	if len(rng) == 0 {
		rang = append(rang, 0, len(bigSlice)) //default is 1 range covering entire input
	} else {
		rang = append(rang, rng...)
	}

	ranges, err := ranger.ValidateAndSortRanges(rang, len(bigSlice))
	if err != nil {
		return bigSlice, rmv, ranges, err
	}
	return bigSlice, rmv, ranges, err
}
