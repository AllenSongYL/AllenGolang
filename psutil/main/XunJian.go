package main

import (
	"bytes"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/thinkeridea/go-extend/exnet"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"
)
// 执行linux命令
func exec_shell(s string) (string, error){
	cmd := exec.Command("/bin/bash","-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// 替换空格回车
func replaceSpace(s string) string {
	var str1 string = ""
	var dsingle uint8 = ','
	var n int = 0
	for _,value :=range s{
		if value!=' ' && value!= '\n'{
			str1+=string(value)
		}else{
			if str1[len(str1) - 1] != dsingle {
				str1+=","
				n += 1
				fmt.Println(n)
				if n == 6{
					str1 = str1 + ";"
					n = 0
					fmt.Println(n)
				}
			}
		}
	}
	return str1
}

func main() {
	// 输出当前时间
	timeNow := time.Now().Format("2006-01-02 15:04:05")

	// 输出内存信息
	memInfo, _ := mem.VirtualMemory()
	memInfoTotal := memInfo.Total/1024/1024/1024
	memInfoUsed := memInfo.Used/1024/1024/1024
	memInfoFree := memInfo.Free/1024/1024/1024
	memInfoPercent := memInfo.UsedPercent
	swapTotal := memInfo.SwapTotal/1024/1024/1024
	swapUsed := memInfo.SwapCached/1024/1024/1024
	swapFree := memInfo.SwapFree/1024/1024/1024

	// 输出CPU信息
	const HunPer float64 = 100
	cpuPer1, _ := cpu.Percent(3*time.Second, false)
	cpuPer2 := cpuPer1[0]
	cpuFreePer := HunPer - cpuPer2

	// CPU负载信息
	cpuUpload, _ := load.Avg()

	// 查看splunk进程
	splCommand := "ps aux | grep -v color | grep splunk | wc -l"
	stdout1, _ := exec_shell(splCommand)
	stdout := strings.TrimSuffix(stdout1, "\n")
	splunklen, _ := strconv.Atoi(stdout)
	var splunkStats string
	if splunklen == 0 {
		splunkStats = "Splunk not running"
	} else {
		splunkStats = "Splunk is running"
	}

	// 获取本地IP
	var ip4addrs []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip4addrs = append(ip4addrs, ipnet.IP.To4())
			}
		}
	}

	var ipDisk map[string]int = map[string]int{
		"182.248.53.42":673,
		"182.248.53.70":354,
		"182.248.53.89":781,
		"182.248.53.90":781,
		"182.248.53.91":781,
		"182.248.53.179":690,
		"182.248.53.180":690,
		"182.248.64.141":788,
		"182.248.64.142":788,
		"182.248.64.143":788,
		"182.248.64.144":793,
		"182.248.64.183":394,
		"182.248.64.184":689,
		"182.248.64.185":689,
		"182.248.64.186":493,
		"182.248.64.187":689,
		"182.248.64.15":576,
		"182.248.64.16":561,
		"182.248.83.158":576,
		"182.248.83.159":576,
		"182.248.64.33":577,
		"182.248.64.34":576,
		"182.248.64.71":576,
		"182.248.53.31":483,
		"182.248.53.45":483,
		"182.248.64.162":185,
		"182.248.64.170":185,
		"182.248.53.41":30,
		"182.248.53.69":390,
		"182.248.64.139":185,
		"182.248.64.151":196,
		"182.248.64.152":185,
		"182.248.64.153":185,
		"182.248.64.154":185,
		"182.248.64.159":185,
		"182.248.64.160":185,
		"182.248.64.145":45,
		"182.248.64.146":45,
		"182.248.64.147":45,
		"182.248.64.148":45,
		"182.248.64.149":45,
		"182.248.64.188":50,
		"182.248.64.189":50,
		"182.248.64.190":50,
		"182.248.64.191":50,
		"182.248.64.192":50,
		"182.248.64.97":168,
		"182.248.64.98":168,
		"182.248.64.99":168,
		"182.248.64.109":168,
		"182.248.64.110":168,
		"182.248.64.111":168,
		"182.248.64.118":168,
		"182.248.64.119":168,
		"182.248.64.120":168,
		"182.248.64.121":168,
		"182.248.64.126":168,
		"182.248.64.127":168,
		"182.248.64.128":168,
		"182.248.64.129":168,
		"182.248.64.132":168,
		"182.248.64.134":168,
		"182.248.64.135":168,
		"182.248.64.136":168,
		"182.248.64.137":168,
		"182.248.64.150":168,
		"182.248.64.161":187,
		"182.248.53.64":344,
		"182.248.53.88":477,
		"182.248.53.11":178,
		"182.248.53.12":178,
		"182.248.53.13":178,
		"182.248.53.14":178,
		"182.248.53.17":178,
		"182.248.53.23":178,
		"182.248.53.25":178,
		"182.248.53.27":178,
		"182.248.64.57":477,
		"182.248.53.87":483,
		"182.248.64.96":576,
		"182.248.57.35":178,
		//"192.168.220.110":101,
	}

	// 获取巡检IP磁盘容量
	var diskInfo string
	for _, values := range ip4addrs {
		ipLong, _ := exnet.IP2Long(values)
		ipString, _ := exnet.Long2IPString(ipLong)
		if _, ok := ipDisk[ipString]; ok {
			diskCommand := "df -h|grep" + " " + strconv.Itoa(ipDisk[ipString])
			diskInfo1, _ := exec_shell(diskCommand)
			diskInfo = replaceSpace(diskInfo1)
			strings.Replace(diskInfo, "\n", ";",-1)
			break
		} else if ipString == "182.248.56.231" {
			diskInfo1, _ := exec_shell("df -h | grep home")
			diskInfo = replaceSpace(diskInfo1)
		}else {
			diskInfo1, _ := exec_shell("df -h | grep -v Avail")
			//fmt.Println(reflect.TypeOf(diskInfo))
			diskInfo = replaceSpace(diskInfo1)
		}
	}

	fmt.Printf("time: %v, ipaddrs: %v, memInfo: %v:%v:%v, swapInfo: %v:%v:%v, " +
		"memPer: %.2f%%, cpuFreePer: %.2f%%, upload: %v, splunkStatus: '%v', diskInfo(Filesystem,Size,Used,Avail,use%%,MountOn): %v\n",
		timeNow,
		ip4addrs,
		memInfoTotal,
		memInfoUsed,
		memInfoFree,
		swapTotal,
		swapUsed,
		swapFree,
		memInfoPercent,
		cpuFreePer,
		cpuUpload,
		splunkStats,
		diskInfo,
		)
}