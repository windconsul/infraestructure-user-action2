package main

import (
	"flag"
	"fmt"
	"infraestructure-user-action/utils"
	"strings"
)

func init() {
	var err error
	flag.StringVar(&domain, "domain", "", "Domain name to be processed")
	flag.StringVar(&instanceName, "instance", "", "Instance name")
	flag.StringVar(&alternativeName, "alternativeName", "", "Alternative name for the instance")
	flag.Parse()

	domain = strings.Trim(domain, " ")
	instanceName = strings.Trim(instanceName, " ")
	alternativeName = strings.Trim(alternativeName, " ")

	if domain == "" {
		fmt.Println("Error: domain is required")
		return
	}
	if instanceName == "" {
		fmt.Println("Error: instance is required")
		return
	}
	user, err = utils.TransformDomainToName(domain)
	if err != nil {
		fmt.Println("Error transforming domain to name:", err)
		return
	}
	if len(user) > 32 && alternativeName == "" {
		fmt.Println("Error: user exceeds 32 characters and no alternative name provided")
		return
	}
	if alternativeName != "" {
		err = utils.ValidateAlternativeDomainName(alternativeName)
		if err != nil {
			fmt.Println("Error validating alternative name:", err)
			return
		}
	}
	if user == "" {
		user = alternativeName
	}
}
