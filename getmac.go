package main

import (
	"fmt"
	"net"
)

const etherType = 0xcccc

func getAllMacAddr() ([]net.HardwareAddr, error) {
	inters, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []net.HardwareAddr
	for _, inter := range inters {
		a := inter.HardwareAddr
		as = append(as, a)
	}
	return as, nil
}

func main() {
	fmt.Println(getAllMacAddr())
}