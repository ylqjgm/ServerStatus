package main

import (
	"github.com/shirou/gopsutil/mem"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
)

func main() {
	v, err := mem.VirtualMemory()
	if nil != err {
		panic(err)
	}
	fmt.Printf("Total: %v, Free: %v, Used: %v, UsedPercent: %f%%\n", v.Total, v.Free, v.Used, v.UsedPercent)
	fmt.Println(v)
	u, err := host.Uptime()
	if nil != err {
		panic(err)
	}
	fmt.Printf("Uptime: %v", u)
	c, err := cpu.Info()
	if nil != err {
		panic(err)
	}
	fmt.Printf("Cpu: %v", c)
	
}
