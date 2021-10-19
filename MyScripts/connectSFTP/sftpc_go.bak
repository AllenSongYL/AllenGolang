package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

// 连接sftp服务器函数
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
	// 测试环境
	auth = append(auth, ssh.Password("111"))

	// 生产环境
	//auth = append(auth, ssh.Password("confbackup"))
	clientConfig = &ssh.ClientConfig{
		// User: SFTP账户名
		User: "allen",
		//User:            "confbackup",
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 测试环境
	addr = fmt.Sprintf("%s:%d", "192.168.220.111", 22)

	// 生产环境
	//addr = fmt.Sprintf("%s:%d", "10.13.132.53", 22)

	sshClient, err = ssh.Dial("tcp", addr, clientConfig)
	if nil != err {
		fmt.Println("ssh 连接失败: ", "10.13.132.53,", err)
	} else {
		fmt.Println("ssh 连接成功: 10.13.132.53 ")
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

// 列出指定远程路径下所有的文件
func listFiles(sftpClient *sftp.Client, remoteFilePath string, localDir string) {
	files, _ := sftpClient.ReadDir(remoteFilePath)
	for _, f := range files {
		if f.IsDir() {
			// 在本地创建同名目录
			fLocalDir := path.Join(localDir, f.Name())

			_, err := os.Stat(fLocalDir)
			if err != nil {
				os.MkdirAll(fLocalDir, os.ModePerm)
				fmt.Println(fLocalDir, ":file create success")
			}
			newRmFile := path.Join(remoteFilePath, f.Name())
			// 如果目录下还有目录则递归调用该函数
			listFiles(sftpClient, newRmFile, fLocalDir)
		} else {
			// 调用getfile() 函数， 下载到本地
			rmfile := path.Join(remoteFilePath, f.Name())
			fmt.Println("Download remote file---> ", rmfile)
			getfile(sftpClient, rmfile, localDir)
		}
	}
}

// 读取指定远程文件，写入到本地
// 传入sftpClient和远程文件名，本地下载路径
func getfile(sftpClient *sftp.Client, rmfile string, localdir string) string {
	// 打开远程文件
	remoteConTest, err := sftpClient.Open(rmfile)
	if err != nil {
		fmt.Println(err)
	}
	// 延迟语句
	// 会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行
	// 先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
	// 放入栈，栈先进后出
	defer remoteConTest.Close()

	var filename = path.Base(rmfile)
	// windows 和 linux 拼接错误
	var fullLocalFile = path.Join(localdir, filename)
	//var fullLocalFile string = localDir + "\\" + localfile

	//fmt.Println("本地文件：" + fullLocalFile)
	downfile, err := os.Create(fullLocalFile)
	if err != nil {
		log.Fatal(err)
	}
	defer downfile.Close()

	if _, err = remoteConTest.WriteTo(downfile); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success download ---> ", fullLocalFile)
	}
	return fullLocalFile
}

func cleanDir(cleandir string, dayago int64) {
	dirInfo, _ := ioutil.ReadDir(cleandir)
	for _, f := range dirInfo {
		fLocalDir := path.Join(cleandir, f.Name())
		if f.IsDir() {
			t, _ := time.ParseInLocation("2006-01-02", f.Name(), time.Local)
			fileUnix := t.Unix()
			if fileUnix <= dayago {
				fmt.Println("删除过期目录 ---> ", fLocalDir)
				os.RemoveAll(fLocalDir)
			}
		}
	}
}

func main() {
	timeStart := time.Now()
	timeStartFormat := timeStart.Format("2006-01-02 15:04:05")
	fmt.Println("程序开始时间：", timeStartFormat)

	// 测试环境
	var remoteFilePath string = "/allen"

	// 生产环境
	//var remoteFilePath string = "/root/confbackup"
	var localDir string = "/opt/data/NetworkConfigBak"

	// 本地目录/opt/data/NetworkConfigBak/2021-09-28
	timeFile := timeStart.Format("2006-01-02")
	localDirData := path.Join(localDir, timeFile)

	// 判断本地路径存不存在，不存在则创建
	_, err := os.Stat(localDirData)
	if err != nil {
		os.MkdirAll(localDirData, os.ModePerm)
		//fmt.Println("file create")
	}

	ftpclient := getConnect()
	fmt.Println("======开始下载文件======")
	listFiles(ftpclient, remoteFilePath, localDirData)

	defer ftpclient.Close()

	// 获取15天前的unix时间戳
	t, _ := time.ParseInLocation("2006-01-02", timeFile, time.Local)
	timeUnix := t.Unix()
	file15Unix := timeUnix - (15 * 24 * 60 * 60)

	fmt.Println("======开始清理过期文件======")
	// 清理本地目录下过期文件
	cleanDir(localDir, file15Unix)

	// 程序结束输出
	timeEnd := time.Now()
	timeEndFormat := timeEnd.Format("2006-01-02 15:04:05")
	fmt.Println("结束时间：", timeEndFormat)
	timeSub := timeEnd.Sub(timeStart)
	fmt.Println("运行时长：", timeSub)
}
