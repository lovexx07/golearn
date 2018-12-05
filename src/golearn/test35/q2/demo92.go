package main

import (
	"time"
	"fmt"
	"net"
)

type dialArgs struct {
	network string
	address string
	timeout time.Duration
}

func main() {
	dialArgList := []dialArgs{
		{
			"tcp",
			"google.cn:80",
			time.Millisecond*500,
		},
		{
			"tcp",
			"google.com:80",
			time.Second*2,
		},

		{
			"tcp",
			"google.com:80",
			time.Minute*4,
		},
	}

	for _, args := range dialArgList{
		fmt.Printf("Dial %q with network %q and timeout %s ...\n", args.address, args.network, args.timeout)

		ts1 := time.Now()
		conn, err := net.DialTimeout(args.network, args.address, args.timeout)
		ts2 := time.Now()
		fmt.Printf("Elapsed time: %s\n", time.Duration(ts2.Sub(ts1)))
		if err != nil{
			fmt.Printf("dial error: %v\n", err)
			fmt.Println()
			continue
		}
		defer  conn.Close()
		fmt.Printf("The local address: %s\n", conn.LocalAddr())
		fmt.Printf("The remove address: %s\n", conn.RemoteAddr())
		fmt.Println()

	}
}
