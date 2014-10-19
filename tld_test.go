package validate

import (
	"strings"
	"testing"
)

func TestIsValidTLD(t *testing.T) {
	for _, tld := range tldList {
		if !IsValidTLD(strings.ToLower(tld)) {
			t.Errorf("%q should be a valid TLD", tld)
		}
		if !IsValidTLD(tld) {
			t.Errorf("%q should be a valid TLD", tld)
		}
	}
	if IsValidTLD("invalid") {
		t.Errorf("%q should not be a valid TLD", "invalid")
	}
}
