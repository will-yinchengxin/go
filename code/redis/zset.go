package utils

import (
	"sort"
)

type ZSet struct {
	scores map[string]float64
}

func NewZSet() *ZSet {
	return &ZSet{
		scores: make(map[string]float64),
	}
}

func (z *ZSet) AddMember(member string, score float64) {
	z.scores[member] = score
}

func (z *ZSet) RemoveMember(member string) {
	delete(z.scores, member)
}

func (z *ZSet) GetScore(member string) (float64, bool) {
	score, exists := z.scores[member]
	return score, exists
}

func (z *ZSet) Members() []string {
	members := make([]string, 0, len(z.scores))
	for member := range z.scores {
		members = append(members, member)
	}
	return members
}

func (z *ZSet) SortMembers() []string {
	members := z.Members()
	sort.Slice(members, func(i, j int) bool {
		return z.scores[members[i]] < z.scores[members[j]]
	})
	return members
}

func (z *ZSet) ZRevRangeByScore(min, max float64) []string {
	return z.SortMembers()
}

// import ( "testing" )
//func TestZSet(t *testing.T) {
//	z := NewZSet()
//	z.AddMember("apple", 1.0)
//	z.AddMember("banana", 2.0)
//	z.AddMember("cherry", 3.0)
//
//	fmt.Println("Members:", z.Members())
//	fmt.Println("Sorted Members:", z.SortMembers())
//
//	z.RemoveMember("banana")
//	fmt.Println("Members after removing 'banana':", z.Members())
//}
