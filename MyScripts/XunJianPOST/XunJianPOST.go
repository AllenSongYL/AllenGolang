package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/parnurzeal/gorequest"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/viper"
	"github.com/thinkeridea/go-extend/exnet"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 执行linux命令
func exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
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
	for _, value := range s {
		if value != ' ' && value != '\n' {
			str1 += string(value)
		} else {
			if str1[len(str1)-1] != dsingle {
				str1 += ","
				n += 1
				if n == 6 {
					str1 = str1[:len(str1)-1] + ";"
					n = 0
				}
			}
		}
	}
	return str1
}

type Iplist struct {
	DiskMap map[string]int `json:"diskmap"`
}

type PostJson struct {
	Ip      string `json:"ip"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func main() {
	// 输出当前时间
	timeStart := time.Now()
	fmt.Println("程序开始运行......")
	timeNow := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("开始时间： ", timeNow)
	filetime := timeStart.Format("20060102_150405")

	log.SetPrefix("[allen]")

	// 初始化发送json结构体
	var postjson PostJson
	postjson.Time = filetime

	// 获取程序当前路径
	ex, _ := os.Executable()
	expath := filepath.Dir(ex)

	//configdir := path.Join(expath, "tarConfig.json")

	// viper 配置读取配置文件
	viper.SetConfigName("iplist")
	viper.SetConfigType("json")
	viper.AddConfigPath(expath)
	fmt.Println("开始读取配置文件--->[iplist.json]")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("读取配置文件失败--->[iplist.json]")
		log.Fatal(err)
	} else {
		log.Println("读取配置文件成功--->[iplist.json]")
	}

	mapdisk := viper.GetStringMap("diskmap")
	//log.Println(mapdisk)
	//for k,v := range mapdisk{
	//	fmt.Printf("mapdisk[%v]: %v",k,v)
	//}

	// 输出内存信息
	memInfo, _ := mem.VirtualMemory()
	memInfoTotal := memInfo.Total / 1024 / 1024 / 1024
	memInfoUsed := memInfo.Used / 1024 / 1024 / 1024
	memInfoFree := memInfo.Free / 1024 / 1024 / 1024
	memInfoPercent := memInfo.UsedPercent
	swapTotal := memInfo.SwapTotal / 1024 / 1024 / 1024
	swapUsed := memInfo.SwapCached / 1024 / 1024 / 1024
	swapFree := memInfo.SwapFree / 1024 / 1024 / 1024

	// 输出CPU信息
	defer ants.Release()

	var cpuFreePer float64
	const HunPer float64 = 100
	cpuPer1, _ := cpu.Percent(3*time.Second, false)
	cpuPer2 := cpuPer1[0]
	cpuFreePer = HunPer - cpuPer2

	// CPU负载信息
	cpuUpload, _ := load.Avg()

	// 查看splunk进程
	splCommand := "ps aux | grep -v grep | grep splunk | wc -l"
	stdout1, _ := exec_shell(splCommand)
	stdout := strings.TrimSuffix(stdout1, "\n")
	splunklen, _ := strconv.Atoi(stdout)
	var splunkStats string
	if splunklen == 0 {
		splunkStats = "Splunk not running"
	} else {
		splunkStats = "Splunk is running"
	}

	// 获取本地所有IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return
	}

	// 拿到巡检IP地址
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ipLong, _ := exnet.IP2Long(ipnet.IP.To4())
			ip4string, _ := exnet.Long2IPString(ipLong)
			if strings.Contains(ip4string, "182.248") {
				log.Println("匹配生产IP地址--->", ip4string)
				postjson.Ip = ip4string
				break
			} else if strings.Contains(ip4string, "192.168") {
				log.Println("匹配测试IP地址--->", ip4string)
				postjson.Ip = ip4string
				break
			}
		}
	}

	// 获取巡检IP磁盘容量
	var diskInfo string

	if _, ok := mapdisk[postjson.Ip]; ok {
		diskdu := fmt.Sprintf("%v", mapdisk[postjson.Ip])
		diskCommand := "df -h | grep " + diskdu
		diskInfo1, _ := exec_shell(diskCommand)
		diskInfo = replaceSpace(diskInfo1)
		strings.Replace(diskInfo, "\n", ";", -1)
	} else if postjson.Ip == "182.248.56.231" {
		diskInfo1, _ := exec_shell("df -h | grep home")
		diskInfo = replaceSpace(diskInfo1)
	} else {
		diskInfo1, _ := exec_shell("df -h | grep -v Avail")
		//fmt.Println(reflect.TypeOf(diskInfo))
		diskInfo = replaceSpace(diskInfo1)
	}

	fresults := fmt.Sprintf("time: %v, ipaddrs: %v, memInfo: %v:%v:%v, swapInfo: %v:%v:%v, "+
		"memPer: %.2f%%, cpuFreePer: %.2f%%, upload: %v, splunkStatus: '%v', diskInfo(Filesystem,Size,Used,Avail,use%%,MountOn): %v ||\n",
		timeNow,
		postjson.Ip,
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
	fresults1 := strings.Replace(fresults, "\"", "", -1)
	fresults2 := strings.Replace(fresults1, "{", "[", -1)
	fresults3 := strings.Replace(fresults2, "}", "]", -1)

	postjson.Message = fresults3

	//rs1, err := json.Marshal(postjson)
	//if err != nil {
	//	fmt.Println("json Marshal err!")
	//}
	//fmt.Printf("发送JSON请求体 ---> %v", string(rs1))

	request := gorequest.New()
	resp, body, errs := request.Post("http://192.168.220.111:8999/post/log").
		Send(postjson).
		End()
	if errs != nil {
		log.Println("发送请求失败...")
	} else {
		log.Println("发送成功...")
	}

	log.Println("发送请求体--->", resp)
	log.Println("返回体--->", body)

	timeEnd := time.Now()
	fmt.Println("程序运行结束......")
	fmt.Println("结束时间： ", timeEnd.Format("2006-01-02 15:04:05"))
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
