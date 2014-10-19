package validate

import (
	"testing"
)

func TestIsValidDomain(t *testing.T) {
	domains := []string{
		"domain.com",
		"domain.net",
		"domain.co.uk",
		"sub.domain.com",
		"sub.domain.co.uk",
	}
	for _, domain := range domains {
		if !IsValidDomain(domain, false) {
			t.Errorf("%q should be a valid domain", domain)
		}
		if IsValidDomain(domain, true) {
			t.Errorf("%q should not be a valid reverse domain", domain)
		}
	}
	reverseDomains := []string{
		"com.domain",
		"net.domain",
		"uk.co.domain",
		"com.domain.sub",
		"uk.co.domain.sub",
	}
	for _, domain := range reverseDomains {
		if !IsValidDomain(domain, true) {
			t.Errorf("%q should be a valid reverse domain", domain)
		}
		if IsValidDomain(domain, false) {
			t.Errorf("%q should not be a valid domain", domain)
		}
	}
}
