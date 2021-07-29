package throttler

import (
	"regexp"
	"strings"
)

// convertSliceToMap convert the list of methods into a map
func convertSliceToMap(s []string) map[string]struct{} {
	mapMethods := make(map[string]struct{})
	for _, v := range s {
		v = strings.ToUpper(v)
		if _, ok := mapMethods[v]; !ok {
			mapMethods[v] = struct{}{}
		}
	}
	return mapMethods
}

// replace the * symbol with +.
func replace(s []string) {
	for i, v := range s {
		s[i] = strings.Replace(v, "*", ".+", -1)
	}
}

// checkPattern check for compliance with the pattern
func checkPattern(prefix string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, _ := regexp.Match(pattern, []byte(prefix))
		if matched {
			return matched
		}
	}
	return false
}
