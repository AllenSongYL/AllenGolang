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
	auth = append(auth, ssh.Password("111")) // 抱歉，这是我电脑密码
	clientConfig = &ssh.ClientConfig{
		// User为账户名
		User:            "allen",
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr = fmt.Sprintf("%s:%d", "192.168.220.111", 22)
	sshClient, err = ssh.Dial("tcp", addr, clientConfig)
	if nil != err {
		fmt.Println("ssh.Dial error", err)
	} else {
		fmt.Println("ssh.Dial 成功")
	}

	// 通过sshClient,创建sftp客户端
	sftpClient, err = sftp.NewClient(sshClient)
	// err不为空，则表示连接失败
	if nil != err {
		fmt.Println("sftp.NewClient error", err)
	} else {
		fmt.Println("sftp.NewClient 成功")
	}
	return sftpClient
}

// 上传文件
//func uploadFile(sftpClient *sftp.Client, localFile, remotePath string) {
//	file, err := os.Open(localFile)
//	if nil != err {
//		fmt.Println("os.Open error", err)
//		return
//	}
//	defer file.Close()
//	remoteFileName := path.Base(localFile)
//	ftpFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName)) // 这里的remotePath是sftp根目录下的目录，是目录不是文件名
//	if nil != err {
//		fmt.Println("sftpClient.Create error", err)
//		return
//	}
//	defer ftpFile.Close()
//	fileByte, err := ioutil.ReadAll(file)
//	if nil != err {
//		fmt.Println("ioutil.ReadAll error", err)
//		return
//	}
//	ftpFile.Write(fileByte)
//}

// 读取远程文件，写入到本地
func getfile(sftpClient *sftp.Client) string {
	var remoteFilePath string = "/allen/a.log"
	var localDir string = "/opt/sftpfile"

	ainfo, err := os.Stat(localDir)
	if err == nil {
		os.MkdirAll(localDir, os.ModePerm)
		fmt.Println("file create")
	} else {
		fmt.Println(ainfo)
	}

	// 打开远程文件
	remoteConTest, err := sftpClient.Open(remoteFilePath)
	if err != nil {
		log.Fatal(err)
	}
	// 延迟语句
	// 会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行
	// 先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。
	// 放入栈，栈先进后出
	defer remoteConTest.Close()

	var localfile string = path.Base(remoteFilePath)
	// windows 和 linux 拼接错误
	var fullLocalFile string = path.Join(localDir, localfile)
	//var fullLocalFile string = localDir + "\\" + localfile

	fmt.Println("本地文件：" + fullLocalFile)
	downfile, err := os.Create(fullLocalFile)
	if err != nil {
		log.Fatal(err)
	}
	defer downfile.Close()

	if _, err = remoteConTest.WriteTo(downfile); err != nil {
		log.Fatal(err)
	}
	return fullLocalFile
}

// 读取,从sftp下载的本地文件
func readFile(lfile string) {
	f, err := os.Open("lfile")
	if err != nil {
		fmt.Println("read file fail", err)
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
	} else {
		fmt.Println(fd)
	}
}

func main() {
	ftpclient := getConnect()
	defer ftpclient.Close()

	localFile := getfile(ftpclient)
	readFile(localFile)
}
