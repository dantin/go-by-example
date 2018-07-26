// version 1.0 2018-07-25

// ex4.5 dedupes a slice of strings.
package main

func unique(strs []string) []string {
	i := 0
	for _, s := range strs {
		if strs[i] == s {
			continue
		}
		i++
		strs[i] = s
	}
	return strs[:i+1]
}
