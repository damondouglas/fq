package set

import "strings"

type StringSet struct {
	internal map[string]struct{}
}

func (s *StringSet) Add(item string) {
	s.internal[item] = struct{}{}
}

func (s *StringSet) Exists(item string) (result bool) {
	_, result = s.internal[item]
	return
}

func (s *StringSet) String() string {
	var tokens []string
	for k := range s.internal {
		tokens = append(tokens, k)
	}
	return strings.Join(tokens, " | ")
}

func NewStringSet(items ...string) (result *StringSet) {
	result = &StringSet{
		internal: map[string]struct{}{},
	}
	for _, k := range items {
		result.internal[k] = struct{}{}
	}
	return
}
