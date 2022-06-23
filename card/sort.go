package deck

type ByRank []Card

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	return a[i].Rank < a[j].Rank
}

type BySuit []Card

func (a BySuit) Len() int      { return len(a) }
func (a BySuit) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySuit) Less(i, j int) bool {
	return a[i].Suit < a[j].Suit
}
