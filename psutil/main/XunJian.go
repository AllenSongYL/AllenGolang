package main

import (
	"bytes"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/sftp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/thinkeridea/go-extend/exnet"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

var sftpUserName = "allen"
var sftpPassword = "111"
var sftpAddress = "192.168.220.111"
var RemoteRootDir = "/allen/XunJian/"

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

// 连接sftp服务器
func getConnect() *sftp.Client {
	// 声明变量类型
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)

	// 创建ssh连接
	auth = make([]ssh.AuthMethod, 0)

	// SFTP账号密码
	auth = append(auth, ssh.Password(sftpPassword))

	clientConfig = &ssh.ClientConfig{
		// User: SFTP账户名
		User:            sftpUserName,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr = fmt.Sprintf("%s:%d", sftpAddress, 22)
	sshClient, err = ssh.Dial("tcp", addr, clientConfig)
	if nil != err {
		fmt.Println("ssh 连接失败: ", sftpAddress, err)
	} else {
		fmt.Println("ssh 连接成功: ", sftpAddress)
	}

	// 通过sshClient,创建sftp客户端
	sftpClient, err = sftp.NewClient(sshClient)
	// err不为空，则表示连接失败
	if nil != err {
		fmt.Println("sftp.NewClient 创建失败", err)
	} else {
		fmt.Println("sftp.NewClient 创建成功")
	}
	return sftpClient
}

// 将获得的结果上传至sftp服务器
func uploadFile(sftpclient *sftp.Client, results []byte, remoteFileDir, filename string) {
	// 判断远程路径是否存在，不存在则创建
	_, err := sftpclient.Stat(remoteFileDir)
	if err != nil {
		errcon := sftpclient.MkdirAll(remoteFileDir)
		if errcon != nil {
			fmt.Println("创建远程路径失败：", remoteFileDir, "。 请检查sftp账号权限！")
			log.Fatal(errcon)
		}
	}

	remoteFullDir := path.Join(remoteFileDir, filename)
	dstfile, err := sftpclient.Create(remoteFullDir)
	if err != nil {
		fmt.Println("创建远程文件失败: ", remoteFullDir)
		log.Fatal(err)
	}
	fmt.Println("成功上传sftp: ", remoteFullDir)
	defer dstfile.Close()

	_, err = dstfile.Write(results)
	if err != nil {
		fmt.Println("写入失败!")
		log.Fatal(err)
	}
}

// 上传目录到sftp服务器
//func uploadDirectory(sftpClient *sftp.Client, localDir, remoteDir string) {
//	Lfile, err := ioutil.ReadDir(localDir)
//	if nil != err {
//		fmt.Println("读取本地目录失败", err)
//		return
//	}
//	// 是目录则新建远程目录，并递归调用进入子目录
//	// 是文件则上传
//	for _, backupDir := range Lfile {
//		if backupDir.IsDir() {
//			localFilePath := path.Join(localDir, backupDir.Name())
//			remoteFilePath := path.Join(remoteDir, backupDir.Name())
//			sftpClient.Mkdir(remoteFilePath)
//			uploadDirectory(sftpClient, localFilePath, remoteFilePath)
//		} else {
//			uploadFile(sftpClient, path.Join(localDir, backupDir.Name()), remoteDir)
//		}
//	}
//}

func main() {
	// 输出当前时间
	timeStart := time.Now()
	fmt.Println("程序开始运行......")
	timeNow := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("开始时间： ", timeNow)
	filetime := timeStart.Format("20060102_150405")

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
	var wg sync.WaitGroup
	var cpuFreePer float64
	wg.Add(1)
	ants.Submit(func() {
		const HunPer float64 = 100
		cpuPer1, _ := cpu.Percent(3*time.Second, false)
		cpuPer2 := cpuPer1[0]
		cpuFreePer = HunPer - cpuPer2
		wg.Done()
	})

	// 连接sftp地址
	var ftpclient *sftp.Client
	wg.Add(1)
	ants.Submit(func() {
		ftpclient = getConnect()
		wg.Done()
	})

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

	// 获取本地IP
	var ip4addrs []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 拿到IP地址
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip4addrs = append(ip4addrs, ipnet.IP.To4())
			}
		}
	}

	var ipDisk map[string]int = map[string]int{
		"182.248.53.42":  673,
		"182.248.53.70":  354,
		"182.248.53.89":  781,
		"182.248.53.90":  781,
		"182.248.53.91":  781,
		"182.248.53.179": 690,
		"182.248.53.180": 690,
		"182.248.64.141": 788,
		"182.248.64.142": 788,
		"182.248.64.143": 788,
		"182.248.64.144": 793,
		"182.248.64.183": 394,
		"182.248.64.184": 689,
		"182.248.64.185": 689,
		"182.248.64.186": 493,
		"182.248.64.187": 689,
		"182.248.64.15":  576,
		"182.248.64.16":  561,
		"182.248.83.158": 576,
		"182.248.83.159": 576,
		"182.248.64.33":  577,
		"182.248.64.34":  576,
		"182.248.64.71":  576,
		"182.248.53.31":  483,
		"182.248.53.45":  483,
		"182.248.64.162": 185,
		"182.248.64.170": 185,
		"182.248.53.41":  30,
		"182.248.53.69":  390,
		"182.248.64.139": 185,
		"182.248.64.151": 196,
		"182.248.64.152": 185,
		"182.248.64.153": 185,
		"182.248.64.154": 185,
		"182.248.64.159": 185,
		"182.248.64.160": 185,
		"182.248.64.145": 45,
		"182.248.64.146": 45,
		"182.248.64.147": 45,
		"182.248.64.148": 45,
		"182.248.64.149": 45,
		"182.248.64.188": 50,
		"182.248.64.189": 50,
		"182.248.64.190": 50,
		"182.248.64.191": 50,
		"182.248.64.192": 50,
		"182.248.64.97":  168,
		"182.248.64.98":  168,
		"182.248.64.99":  168,
		"182.248.64.109": 168,
		"182.248.64.110": 168,
		"182.248.64.111": 168,
		"182.248.64.118": 168,
		"182.248.64.119": 168,
		"182.248.64.120": 168,
		"182.248.64.121": 168,
		"182.248.64.126": 168,
		"182.248.64.127": 168,
		"182.248.64.128": 168,
		"182.248.64.129": 168,
		"182.248.64.132": 168,
		"182.248.64.134": 168,
		"182.248.64.135": 168,
		"182.248.64.136": 168,
		"182.248.64.137": 168,
		"182.248.64.150": 168,
		"182.248.64.161": 187,
		"182.248.53.64":  344,
		"182.248.53.88":  477,
		"182.248.53.11":  178,
		"182.248.53.12":  178,
		"182.248.53.13":  178,
		"182.248.53.14":  178,
		"182.248.53.17":  178,
		"182.248.53.23":  178,
		"182.248.53.25":  178,
		"182.248.53.27":  178,
		"182.248.64.57":  477,
		"182.248.53.87":  483,
		"182.248.64.96":  576,
		"182.248.57.35":  178,
		//"192.168.220.110":101,
	}

	// 获取巡检IP磁盘容量
	var diskInfo string
	wg.Add(1)
	ants.Submit(func() {
		for _, values := range ip4addrs {
			ipLong, _ := exnet.IP2Long(values)
			ipString, _ := exnet.Long2IPString(ipLong)
			if _, ok := ipDisk[ipString]; ok {
				diskCommand := "df -h|grep" + " " + strconv.Itoa(ipDisk[ipString])
				diskInfo1, _ := exec_shell(diskCommand)
				diskInfo = replaceSpace(diskInfo1)
				strings.Replace(diskInfo, "\n", ";", -1)
				break
			} else if ipString == "182.248.56.231" {
				diskInfo1, _ := exec_shell("df -h | grep home")
				diskInfo = replaceSpace(diskInfo1)
			} else {
				diskInfo1, _ := exec_shell("df -h | grep -v Avail")
				//fmt.Println(reflect.TypeOf(diskInfo))
				diskInfo = replaceSpace(diskInfo1)
			}
		}
		wg.Done()
	})

	wg.Wait()

	fresults := fmt.Sprintf("time: %v, ipaddrs: %v, memInfo: %v:%v:%v, swapInfo: %v:%v:%v, "+
		"memPer: %.2f%%, cpuFreePer: %.2f%%, upload: %v, splunkStatus: '%v', diskInfo(Filesystem,Size,Used,Avail,use%%,MountOn): %v ||\n",
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

	results := []byte(fresults)

	// 上传至sftp
	// 测试环境地址
	var remoteFilePath = path.Join(RemoteRootDir, ip4addrs[0].String())
	uploadFile(ftpclient, results, remoteFilePath, filetime+".log")

	timeEnd := time.Now()
	fmt.Println("程序运行结束。")
	fmt.Println("结束时间： ", timeEnd.Format("2006-01-02 15:04:05"))
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
