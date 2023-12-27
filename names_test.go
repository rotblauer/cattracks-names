package names

import (
	"testing"
)

// TestAliasesToMatchers_Unique tests that all known aliases (eg. 'rye', 'ia') are unique.
func TestAliasesToMatchers_Unique(t *testing.T) {
	index := make(map[string]bool, len(AliasesToMatchers))
	for alias := range AliasesToMatchers {
		if _, ok := index[alias]; !ok {
			index[alias] = true
		} else {
			t.Errorf("alias %s is not unique", alias)
		}
	}
}
