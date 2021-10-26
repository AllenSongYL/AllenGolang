package main

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
)

func main() {
	// 磁盘IO统计
	a, _ := net.IOCounters(true)
	//fmt.Println(a)
	for k, v := range a {
		fmt.Println(k)
		fmt.Println(v)
	}

}
