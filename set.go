package shortuuid

import (
	"sort"
	"strings"
)

type StringSet struct {
	set map[string]bool
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}

func (set *StringSet) Add(i string) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *StringSet) Contains(i string) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *StringSet) Remove(i string) {
	delete(set.set, i)
}

func (set *StringSet) Len() int {
	return len(set.set)
}

func (set *StringSet) String() string {
	var strs []string
	for s, _ := range set.set {
		strs = append(strs, s)
	}

	ss := sort.StringSlice(strs)
	sort.Sort(ss)

	return strings.Join(ss, "")
}
