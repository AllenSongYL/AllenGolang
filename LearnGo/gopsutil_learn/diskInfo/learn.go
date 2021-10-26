package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"reflect"
)

func main() {
	// 磁盘IO统计
	mapstats, err := disk.IOCounters()
	if err != nil {
		fmt.Println("获取磁盘信息失败！！！")
	}
	for name, state := range mapstats {
		fmt.Printf("name: %v, state: %v\n", name, state)

	}

	// 磁盘分区信息  返回值：device 分区标识；mountpoint 挂载点；fstype 文件系统类型；opts 选项，与系统有关
	diskPartition, _ := disk.Partitions(true)
	fmt.Println("disk--->", diskPartition)

	// 磁盘使用率
	diskuse, _ := disk.Usage("C:")
	fmt.Println("C盘使用情况---> ", diskuse)
	fmt.Println(reflect.TypeOf(diskuse))
	diskuseData, _ := json.MarshalIndent(diskuse, "", " ")
	fmt.Println(string(diskuseData))
}
