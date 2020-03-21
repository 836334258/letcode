package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Replace(address,".","[.]",-1)
}

func main()  {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))
}
