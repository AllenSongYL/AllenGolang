package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	// 测试环境
	centosCommand := "curl -k -u admin:'1qaz!QAZ' http://192.168.220.110:8089/services/data/inputs/udp/514 " +
		"-d \"restrictToHost=" + ips + "&index=crcb_unix&sourcetype=centos_unix\""
	// 生产环境命令
	//centosCommand := "curl -k -u admin:'1qaz!QAZ' https://170.130.106.23:8089/services/data/inputs/udp/514 " +
	//	"-d \"restrictToHost=" + ips + "&index=crcb_unix&sourcetype=centos_unix\""

	if ipType == "CENTOS" {
		results, err := exec_shell(centosCommand)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(results)
		}
	}

}

func main() {
	// 打开csv文件
	csvName := "G:\\GO\\笔记\\AllenGolang\\CurlAddInputs\\csvDir\\ip_assets.csv"
	csvFile, err := os.Open(csvName)
	if err != nil {
		log.Fatal("无法打开该文件", err)
	}
	defer csvFile.Close()

	//
	r := csv.NewReader(csvFile)

	record, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, values := range record {
		exec_curl(values[0], values[1])
	}
}
