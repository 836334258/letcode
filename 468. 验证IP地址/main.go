package main

import (
	"fmt"
	"regexp"
)

func validIPAddress(IP string) string {
	matchString, _ := regexp.MatchString("[1-9][0-9]{3}.[1-9][0-9]{3}.[1-9][0-9]{3}.[1-9][0-9]{3}", IP)
	fmt.Println(matchString)

	return "IPv4"
}

func main() {
	IP := "172.16.254.1"
	fmt.Println(validIPAddress(IP))
}
