package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func cleanDir(cleandir string, today int64, dayago int64) {
	err := os.Chdir(cleandir)
	if err != nil {
		fmt.Printf("切换到目录： %v 失败\n", cleanDir)
	} else {
		dirInfo, _ := ioutil.ReadDir(cleandir)
		for _, f := range dirInfo {
			fLocalDir := path.Join(cleandir, f.Name())
			if f.IsDir() {
				// 递归： 移动到子目录下继续
				cleanDir(fLocalDir, today, dayago)
			} else {
				fielChangeTime := f.ModTime()
				filechangeUnix := fielChangeTime.Unix()
				chaTime := today - filechangeUnix
				if chaTime > dayago {
					os.Remove(fLocalDir)
					fmt.Println("删除过期文件: ", fLocalDir, ",变更时间: ", fielChangeTime.Format("2006-01-02 15:04:05"))
				}
			}
		}
	}
}

func main() {

	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	//
	timeStartFormat2 := timeStart.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStartFormat2, time.Local)
	timeUnix := t.Unix()
	fmt.Println(timeUnix)
	file15Unix := timeUnix - (15 * 24 * 60 * 60)
	fmt.Println(file15Unix)
	//

	timeStartUnix := timeStart.Unix()
	fmt.Println("开始时间：", timeStartFormat)

	// 清理15天之前的文件
	var dagAgo int64
	dagAgo = 15 * 24 * 60 * 60

	// 清理目录 G:\GO\测试环境
	var cDir string = "G:\\GO\\测试环境"
	cleanDir(cDir, timeStartUnix, dagAgo)

	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)

}
