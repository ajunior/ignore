package main

type set map[string]struct{}

func (s set) add(template string) {
	s[template] = struct{}{}
}

func (s set) has(template string) bool {
	_, exists := s[template]
	return exists
}
