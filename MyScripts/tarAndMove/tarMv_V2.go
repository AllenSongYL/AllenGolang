package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// 遍历指定的目录，遇到目录则递归
// soucedir:  "/home/applog/531PBANK/ASCPEBS201"
// destdir: "/home/applog/531PBANK_1/ASCPEBS201"
func forDir(sourcedirs, destdir string) {
	listfile, _ := ioutil.ReadDir(sourcedirs)
	for _, file := range listfile {
		if file.IsDir() {
			destdir2 := path.Join(destdir, file.Name())
			if _, err := os.Stat(destdir2); err != nil {
				os.MkdirAll(destdir2, 644)
			}
			forDir(path.Join(sourcedirs, file.Name()), destdir2)
		} else if strings.Contains(file.Name(), ".tar.Z") {
			fullfile := path.Join(sourcedirs, file.Name())
			_, err := exec_shell("tar -Zxvf " + fullfile + " -C " + destdir)
			if err != nil {
				log.Println("打开压缩文件失败：", err)
			}
			log.Printf("<--- 解压完成: %v --->", fullfile)
		}
	}
}

func main() {
	timeStart := time.Now()
	fmt.Println("程序开始运行......")
	timeNow := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("开始时间： ", timeNow)

	// 获取程序当前路径
	ex, _ := os.Executable()
	expath := filepath.Dir(ex)

	//configdir := path.Join(expath, "tarConfig.json")

	// viper 配置读取配置文件
	viper.SetConfigName("tarConfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(expath)
	fmt.Println("读取配置文件--->[tarConfig.json]")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败！！！")
		log.Fatal(err)
	} else {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), ":  读取成功！")
	}

	// 获取配置文件中的rootDir
	rootDir := viper.GetStringSlice("rootDir")
	targetDir := viper.GetString("targetDir")

	for _, i := range rootDir {
		filename := filepath.Base(i)
		targetDir2 := path.Join(targetDir, filename)
		if _, err := os.Stat(targetDir2); err != nil {
			os.MkdirAll(targetDir2, 0644)
		}
		forDir(i, targetDir2)
	}

	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)

}
