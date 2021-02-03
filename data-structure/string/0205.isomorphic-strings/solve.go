package problem0205

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := map[byte]byte{}
	set := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		m, found := dict[s[i]]
		if found && m != t[i] {
			return false
		}
		if !found {
			if _, setFound := set[t[i]]; setFound {
				return false
			}
			dict[s[i]] = t[i]
			set[t[i]] = s[i]
		}
	}
	return true
}
