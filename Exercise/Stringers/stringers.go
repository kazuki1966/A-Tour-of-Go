package main

import (
	"fmt"
	"strings"
)

type ipAddr [4]byte

func (ip ipAddr) String() string {
	ipStr := make([]string, 4)

	for i, v := range ip {
		ipStr[i] = fmt.Sprintf("%v", v)
	}

	return strings.Join(ipStr, ".")
}

func main() {
	hosts := map[string]ipAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
