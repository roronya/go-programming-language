package main

func f(s []string) []string {
	if len(s) == 0 {
		return s
	}
	out := s[:1]
	for _, v := range s {
		if out[len(out)-1] == v {
			continue
		}
		out = append(out, v)
	}
	return out
}
