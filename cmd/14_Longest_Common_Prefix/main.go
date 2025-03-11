package main

func longestCommonPrefix(strs []string) string {
	r := []byte{}

	for i := range 200 {
		if i == len(strs[0]) {
			break
		}

		for j := range len(strs) - 1 {
			if i == len(strs[j+1]) {
				return string(r)
			}

			if strs[0][i] != strs[j+1][i] {
				return string(r)
			}
		}

		r = append(r, strs[0][i])
	}

	return string(r)
}
