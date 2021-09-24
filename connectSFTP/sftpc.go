package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

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
	auth = append(auth, ssh.Password("confbackup")) // 抱歉，这是我电脑密码
	clientConfig = &ssh.ClientConfig{
		// User为账户名
		User:            "confbackup",
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr = fmt.Sprintf("%s:%d", "192.168.220.111", 22)
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
		log.Fatal(err)
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
		log.Fatal(err)
	} else {
		fmt.Println("success download ---> ", fullLocalFile)
	}
	return fullLocalFile
}

// 读取,从sftp下载的本地文件
func readFile(lfile string) {
	f, _ := os.Open(lfile)
	defer f.Close()

	fd, _ := ioutil.ReadAll(f)
	result := strings.Replace(string(fd), "\n", "", 1)
	fmt.Println(result)
}

func main() {
	var remoteFilePath string = "/root/confbackup"
	var localDir string = "/opt/splunk/data"

	// 判断本地路径存不存在，不存在则创建
	_, err := os.Stat(localDir)
	if err != nil {
		os.MkdirAll(localDir, os.ModePerm)
		//fmt.Println("file create")
	}

	ftpclient := getConnect()
	listFiles(ftpclient, remoteFilePath, localDir)

	defer ftpclient.Close()
}
