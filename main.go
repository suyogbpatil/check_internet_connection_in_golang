package main

import (
	"fmt"
	"net"
)

func checkinternetconnection() bool {
	addresses, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, address := range addresses {

		flagup := address.Flags&net.FlagUp != 0
		loopbackup := address.Flags&net.FlagLoopback != 0

		if !loopbackup {
			fmt.Printf("Name %s \nConnected Status %v  \nFlagLoopback %v \n", address.Name, flagup, loopbackup)
			fmt.Println("----")
			if flagup {
				count++
			}

		}

	}
	return count != 0

}

func main() {

	internetcheck := checkinternetconnection()
	if internetcheck {
		fmt.Println("Check Connection is Connected ")
	} else {
		fmt.Println("Internet Connection Disconnected")
	}

}
