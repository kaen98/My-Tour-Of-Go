package main

import (
	"fmt"
	"strings"
	"strconv"
)


type IPAddr [4]byte

func (I IPAddr) String() string {
	var ip_list = make([]string, len(I))
	for i, v := range I {
		ip_list[i] = strconv.Itoa(int(v))
	}
	return strings.Join(ip_list, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}