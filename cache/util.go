package cache

import "strings"

const (
	helixQueryLimit = 100
)

// Remove duplicates from a string slice (case insensitive)
func removeDuplicates(slice []string) []string {
	m := make(map[string]struct{}, len(slice))
	var out []string
	for _, val := range slice {
		val = strings.ToLower(val)
		if _, has := m[val]; !has {
			m[val] = struct{}{}
			out = append(out, val)
		}
	}
	return out
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateQuerySliceIndexes(a, b []string) (int, int) {
	if len(a) < helixQueryLimit {
		return len(a), min(helixQueryLimit-len(a), len(b))
	} else {
		return helixQueryLimit, 0
	}
}

func formatString(prefix, input string) string {
	return prefix + strings.ToLower(input)
}
