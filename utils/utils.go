package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateAlternativeDomainName(name string) (err error) {
	if len(name) > 32 {
		return fmt.Errorf("alternative name exceeds 32 characters")
	}

	validName := regexp.MustCompile(`^[a-z_][a-z0-9_-]*[$]?$`)
	if !validName.MatchString(name) {
		return fmt.Errorf("invalid alternative name format")
	}

	return
}
func TransformDomainToName(domain string) (name string, err error) {

	name = strings.ReplaceAll(domain, ".", "_")

	validName := regexp.MustCompile(`^[a-z_][a-z0-9_-]*[$]?$`)
	if !validName.MatchString(name) {
		return "", fmt.Errorf("invalid domain name format")
	}

	return
}
