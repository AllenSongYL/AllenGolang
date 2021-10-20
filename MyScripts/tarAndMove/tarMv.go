package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// 遍历指定的目录，遇到目录则递归
// soucedir:  "/home/applog/531PBANK/ASCPEBS201"
// destdir: "/home/applog/531PBANK_1/ASCPEBS201"
func forDir(sourcedirs, destdir string) {
	listfile, _ := ioutil.ReadDir(sourcedirs)
	fmt.Println(listfile)
	for _, file := range listfile {
		if file.IsDir() {
			destdir2 := path.Join(destdir, file.Name())
			fmt.Println(destdir2)
			if _, err := os.Stat(destdir2); err != nil {
				os.MkdirAll(destdir2, 644)
			}
			forDir(path.Join(sourcedirs, file.Name()), destdir2)
		} else if strings.Contains(file.Name(), ".tar") {
			fmt.Println("匹配到tar包", file.Name())
			// 打开源tar包
			tarfile, err := os.Open(path.Join(sourcedirs, file.Name()))
			if err != nil {
				fmt.Println("打开压缩包失败！！！")
				log.Fatal(err)
			} else {
				fmt.Println("开始解压------>", path.Join(sourcedirs, file.Name()))
			}

			defer tarfile.Close()
			// 将打开的文件解压
			untarfile, err := gzip.NewReader(tarfile)
			if err != nil {
				log.Fatal(err)
			}
			defer untarfile.Close()
			// 创建tar.Reader 结构
			untarfile2 := tar.NewReader(untarfile)

			for {
				f, err := untarfile2.Next()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}

				dstFileOrDie := path.Join(destdir, f.Name)
				fmt.Println(dstFileOrDie)

				switch f.Typeflag {
				// 是文件则写入
				case tar.TypeReg:
					fw, err := os.OpenFile(dstFileOrDie, os.O_CREATE|os.O_WRONLY, 0644)
					fmt.Println("解压文件------>", dstFileOrDie)
					if err != nil {
						log.Fatal(err)
					}
					defer fw.Close()

					_, err = io.Copy(fw, untarfile2)
					if err != nil {
						fmt.Println("解压写入失败！！！")
					} else {
						fmt.Println("<--- 解压成功 --->")
					}

				// 是目录则创建
				case tar.TypeDir:
					if _, err := os.Stat(dstFileOrDie); err != nil {
						if err := os.MkdirAll(dstFileOrDie, 0644); err != nil {
							fmt.Println("创建解压目录失败！！！")
							log.Fatal(err)
						}
					}

				}
			}
		}
	}
}

var wg sync.WaitGroup

func main() {
	timeStart := time.Now()
	fmt.Println("程序开始运行......")
	timeNow := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("开始时间： ", timeNow)

	// 获取程序当前路径
	ex, _ := os.Executable()
	expath := filepath.Dir(ex)

	//configdir := path.Join(expath, "tarConfig.json")
	defer ants.Release()
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
		wg.Add(1)

		ants.Submit(func() {
			forDir(i, targetDir2)
			wg.Done()
		})
	}
	wg.Wait()
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)

}
