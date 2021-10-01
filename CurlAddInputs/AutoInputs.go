package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func exec_curl(ips, ipType string) {

	// own测试环境
	//CurlCommand := "curl -k -u admin:'1qaz!QAZ' https://192.168.220.110:8089/services/data/inputs/udp/514 " +
	//	"-d \"restrictToHost=" + ips + "&index=crcb_unix&sourcetype="

	// 农商行测试环境
	//CurlCommand := "curl -k -u admin:'1qaz!QAZ' https://170.130.106.23:8089/services/data/inputs/udp/514 " +
	//	"-d \"restrictToHost=" + ips + "&index=crcb_unix&sourcetype="

	// 生产环境命令
	CurlCommand := "curl -k -u input:'1234qwer' https://10.13.130.21:8089/services/data/inputs/udp/514 " +
		"-d \"restrictToHost=" + ips + "&index=crcb_unix&sourcetype="

	if ipType == "CENTOS" {
		fmt.Println("开始添加centos_unix：", ips)
		_, err := exec_shell(CurlCommand + "centos_unix\"")
		if err != nil {
			log.Fatal("执行curl命令失败：", err)
		} else {
			fmt.Println("添加成功！")
		}
	} else if ipType == "REDHAT" {
		fmt.Println("开始添加redhat_unix：", ips)
		_, err := exec_shell(CurlCommand + "redhat_unix\"")
		if err != nil {
			log.Fatal("执行curl命令失败：", err)
		} else {
			fmt.Println("添加成功！")
		}
	} else if ipType == "AIX" {
		fmt.Println("开始添加aix_unix：", ips)
		_, err := exec_shell(CurlCommand + "aix_unix\"")
		if err != nil {
			log.Fatal("执行curl命令失败：", err)
		} else {
			fmt.Println("添加成功！")
		}
	} else {
		fmt.Println("开始添加suse_unix：", ips)
		_, err := exec_shell(CurlCommand + "suse_unix\"")
		if err != nil {
			log.Fatal("执行curl命令失败：", err)
		} else {
			fmt.Println("添加成功！")
		}
	}
}

func main() {

	// 测试环境 打开csv文件
	//csvName := "G:\\GO\\笔记\\AllenGolang\\CurlAddInputs\\csvDir\\ip_assets.csv"
	//csvName := "/opt/splunk/etc/apps/search/lookups/test_input.csv"

	// 生产环境
	csvName := "/data/splunk/etc/apps/crcb_soc/lookups/auto_udp_inputs.csv"

	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	csvFile, err := os.Open(csvName)
	if err != nil {
		log.Fatal("无法打开该文件", err)
	} else {
		fmt.Println("成功打开文件：", csvName)
	}
	defer csvFile.Close()

	//
	r := csv.NewReader(csvFile)

	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("成功解析文件：", csvName)
	}

	for _, values := range record {
		exec_curl(values[0], values[1])
	}

	// 程序结束输出
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
