package rmvtag

import (
	"fmt"
	"sort"
)

// validateAndSortRanges ensures ranges entered on the command line are sane and usable.
// basic validation means ensuring even number of integers, start < end, no overlap,
// start < last line and delivered to callers in ascending order
func validateAndSortRanges(r []int, max int) ([][2]int, error) {
	var pair [2]int
	var pairs [][2]int

	if len(r) == 0 {
		r[0] = 0
		r[1] = max
	}

	if len(r) < 2 { //must be at least one pair
		return pairs, fmt.Errorf("ranges passed on command line must be at least one pair, got %v", r)

	}

	if (len(r) % 2) != 0 { // must be even number
		return pairs, fmt.Errorf("ranges passed on command line must be even number of integers, got %v", r)
	}

	for i := 0; i < len(r); i += 2 { //arrange into pairs
		pair[0] = r[i]
		pair[1] = r[i+1]
		pairs = append(pairs, pair)
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i][0] < pairs[j][0] })

	for i := 0; i < len(pairs); i++ {
		if pairs[i][0] >= pairs[i][1] { //range start must be < end
			return [][2]int{}, fmt.Errorf("range passed has start %d >= end %d", pairs[i][0], pairs[i][1])
		}
	}

	if len(pairs) > 1 {
		for i := 0; i < len(pairs)-1; i++ {
			if pairs[i][1] > pairs[i+1][0] { //ranges must not overlap
				return [][2]int{}, fmt.Errorf("ranges passed on command line cannot overlap")
			}
		}
	}

	// if len(pairs) >= 2 {
	// 	if pairs[len(pairs)-2][1] > pairs[len(pairs)-1][0] { //check last one
	// 		return [][2]int{}, fmt.Errorf("ranges passed on command line cannot overlap")
	// 	}
	// }

	for i := 0; i < len(pairs); i++ { // make sure a start is not past eof
		if pairs[i][0] > max {
			return [][2]int{}, fmt.Errorf("range passed on command line %d %d is past end of input", pairs[i][1], pairs[i][1])
		}
	}

	return pairs, nil
}
