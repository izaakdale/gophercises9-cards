// Code generated by "stringer -type=RankType"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[JokerT-0]
	_ = x[Ace-1]
	_ = x[Two-2]
	_ = x[Three-3]
	_ = x[Four-4]
	_ = x[Five-5]
	_ = x[Six-6]
	_ = x[Seven-7]
	_ = x[Eight-8]
	_ = x[Nine-9]
	_ = x[Ten-10]
	_ = x[Jack-11]
	_ = x[Queen-12]
	_ = x[King-13]
}

const _RankType_name = "JokerTAceTwoThreeFourFiveSixSevenEightNineTenJackQueenKing"

var _RankType_index = [...]uint8{0, 6, 9, 12, 17, 21, 25, 28, 33, 38, 42, 45, 49, 54, 58}

func (i RankType) String() string {
	if i < 0 || i >= RankType(len(_RankType_index)-1) {
		return "RankType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RankType_name[_RankType_index[i]:_RankType_index[i+1]]
}
