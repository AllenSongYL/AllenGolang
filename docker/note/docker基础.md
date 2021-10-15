# 概念介绍

## 镜像

镜像是一种轻量级，可执行的独立软件包，用来打包软件运行环境和基本运行环境开发的软件，它包含运行某个软件所需的所有内容，包括代码，运行时，库，环境变量和配置文件。

Docker镜像(Image)就是一个只读的模板。镜像可以用来创建Docker容器，一个镜像可以创建多个容器。



## 容器

容器是用镜像创建的运行实例。

可以被启动，停止，删除。每个容器都是相互隔离的。保证安全的平台。

当容器启动时，一个新的可写层被加载到镜像的顶部。

这一层通常被称作容器层，容器层之下的都叫镜像层。

镜像和容器的关系类似于面向对象编程中的类和对象。

| Docker | 面向对象 |
| ------ | -------- |
| 镜像   | 类       |
| 容器   | 对象     |





## UnionFS 联合文件系统

是一种分层，轻量级并且高性能的文件系统，它支持堆文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统。





## 下载镜像

docker  pull  mongo

# 命令

## docker  version  

// docker版本



## docker  info        

// docker 详细信息



## docker --help        

// 查看命令帮助 Usage:  docker [OPTIONS] COMMAND

 

## docker images    

// 列出镜像模板

​		参数： - a  所有镜像（含中间映像层）

​					-q   只显示镜像ID

​                    --digests   显示镜像的摘要信息

​				  --no-trunc   显示完整的镜像信息



~~~
PS C:\Users\alan> docker images
REPOSITORY    TAG       IMAGE ID       CREATED        SIZE
mysql         latest    9da615fced53   11 hours ago   514MB
golang        latest    8e64853afe20   5 days ago     941MB
python        latest    618fff2bfc18   7 days ago     915MB
hello-world   latest    feb5d9fea6a5   2 weeks ago    13.3kB
centos        centos7   eeb6ee3f44bd   3 weeks ago    204MB
~~~

​		REPOSITORY： 表示镜像的仓库源

​		TAG： 镜像ID

​		CREATED: 镜像创建时间

​		SIZE:  镜像大小

​		同一个仓库源可以有多个TAG，代表这个仓库源的不同版本，我们使用
​        REPOSITORY:TAG来定义不同的镜像。如果不指定一个镜像的版本标签，
​        默认使用 "xxx:latest" 镜像



## docker search

// 搜索镜像

docker search --filter=stars=30 tomcat    

// 筛选出stars数量超过30的tomcat镜像

~~~
PS C:\Users\alan> docker search tomcat
NAME                          DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
tomcat                        Apache Tomcat is an open source implementati…   3150      [OK]
tomee                         Apache TomEE is an all-Apache Java EE certif…   93        [OK]
dordoka/tomcat                Ubuntu 14.04, Oracle JDK 8 and Tomcat 8 base…   58                   [OK]
kubeguide/tomcat-app          Tomcat image for Chapter 1                      31
consol/tomcat-7.0             Tomcat 7.0.57, 8080, "admin/admin"              18                   [OK]
cloudesire/tomcat             Tomcat server, 6/7/8                            15                   [OK]
aallam/tomcat-mysql           Debian, Oracle JDK, Tomcat & MySQL              13                   [OK]
arm32v7/tomcat                Apache Tomcat is an open source implementati…   11
rightctrl/tomcat              CentOS , Oracle Java, tomcat application ssl…   7                    [OK]
arm64v8/tomcat                Apache Tomcat is an open source implementati…   6
maluuba/tomcat7-java8         Tomcat7 with java8.                             6
unidata/tomcat-docker         Security-hardened Tomcat Docker container.      5                    [OK]
amd64/tomcat                  Apache Tomcat is an open source implementati…   3
fabric8/tomcat-8              Fabric8 Tomcat 8 Image                          2                    [OK]
jelastic/tomcat               An image of the Tomcat Java application serv…   2
oobsri/tomcat8                Testing CI Jobs with different names.           2
cfje/tomcat-resource          Tomcat Concourse Resource                       2
picoded/tomcat7               tomcat7 with jre8 and MANAGER_USER / MANAGER…   1                    [OK]
ppc64le/tomcat                Apache Tomcat is an open source implementati…   1
chenyufeng/tomcat-centos      tomcat基于centos6的镜像                              1                    [OK]
99taxis/tomcat7               Tomcat7                                         1                    [OK]
camptocamp/tomcat-logback     Docker image for tomcat with logback integra…   1                    [OK]
secoresearch/tomcat-varnish   Tomcat and Varnish 5.0                          0                    [OK]
softwareplant/tomcat          Tomcat images for jira-cloud testing            0                    [OK]
s390x/tomcat                  Apache Tomcat is an open source implementati…   0
~~~



## docker rmi  

// 删除镜像

docker  rmi  -f  $(docker  images  -q)



## 容器命令

### docker run 镜像名:标签



docker  run  hello-world

docker run -it --name="centos_v1"  centos:centos7

参数：

- --name="容器新名字"       //  为容器指定一个名称

-  -d                                    //   后台运行容器，并返回容器ID，也即启动守护式容器

  ~~~
  docker ps 
  无法看到
  ~~~

  

- -i                                      //  以交互模式运行容器，通常与-t同时使用

- -t                                      //  为容器重新分配一个伪输入终端，通常与-i同时使用

- -P                                    //   随机端口映射

- -p                                   //   指定端口映射，以下四种格式

  1. ip:

![docker_run运行过程](G:\GO\笔记\AllenGolang\docker\note\图\docker_run运行过程.png)



### docker   ps

参数：

1. -q     // 只显示容器编号
2. -a     // 显示所有容器

//  列出所有正在运行的容器

~~~
PS C:\Users\alan> docker ps
CONTAINER ID   IMAGE            COMMAND       CREATED          STATUS          PORTS     NAMES
df852c642704   centos:centos7   "/bin/bash"   26 seconds ago   Up 26 seconds             centos_v1
~~~



## 退出容器

#### exit

容器停止并退出

#### ctrl + P + Q

容器不停止退出



## 启动容器

docker  start   容器ID



## 重启容器

docker  restart  容器ID



## 停止容器

docker stop  容器ID



## 强制关闭容器

docker  kill  容器ID



## 删除已停止的容器

docker   rm   容器ID

  强制删除使用 -f



## 查看容器日志

docker log -f -t --tail 10 容器ID

参数

1.  -t    // 加入时间戳
2. -f    //  跟随最新的日志打印
3. --tail  数字  // 显示最后多少条

~~~
docker run -d centos:centos7 /bin/sh -c "while true;do echo hello syl;sleep 2;done"

PS C:\Users\alan> docker logs -t -f --tail 100 6bc7555e13b0
2021-10-13T05:05:48.828019736Z hello syl
2021-10-13T05:05:50.829131069Z hello syl
2021-10-13T05:05:52.830145371Z hello syl
2021-10-13T05:05:54.831011061Z hello syl
2021-10-13T05:05:56.832020066Z hello syl
2021-10-13T05:05:58.833066269Z hello syl
2021-10-13T05:06:00.834033919Z hello syl
2021-10-13T05:06:02.834900590Z hello syl
2021-10-13T05:06:04.835996909Z hello syl
2021-10-13T05:06:06.836921894Z hello syl
2021-10-13T05:06:08.837944039Z hello syl
2021-10-13T05:06:10.839042437Z hello syl
~~~



## 查看容器内运行的进程

docker  top  容器ID



## 查看容器内部的细节

docker  inspect   容器ID



## 进入正在运行的容器并以命令行交互

docker   attach  容器ID

直接进入容器启动命令终端，不会启动新的进程



**docker  exec  -it 容器ID  ls -l  /tmp**

在容器外执行操作

实在容器中打开新的终端，并且可以启动新的进程

**docker  exec  -it 容器ID   /bin/bash**

进入容器



## 从容器内拷贝文件

docker cp 容器ID:/tmp/yum.log  /root



## 容器提交

docker  commit  -m="xx" -a="xx" 容器ID   目标镜像名:标签名

- -m="xxx" 指定提交的描述信息
- -a="xxx"  指定作者

提交容器副本使之称为一个新的镜像

~~~
docker commit -a="syl" -m="without docs"  容器ID  axxx/tomcat:1.2
~~~

# 容器数据卷



### 直接命令添加

**docker run -it -v /宿主机绝对路径:/容器内路径  (--privileged=true) 镜像名**

没有会创建

宿主机和容器之间数据共享。

容器停止退出后，主机修改数据时，容器也能同步。

### 命令i添加加上权限

**docker run -it -v /宿主机绝对路径:/容器内路径:ro  镜像名**

//  只读

### dockerfile添加





## 数据卷容器

命名的容器挂载数据卷，其他容器通过挂载这个（父容器）实现数据共享，挂载数据卷的容器，称之为数据卷容器

dc02的容器继承dc01的卷

~~~
docker run -it --name dc02 --volumes-from dc01 zzyy/centos
~~~

父子各自添加都能共享。



# Docker 网络

## 容器地址

~~~
[root@df852c642704 /]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: tunl0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/ipip 0.0.0.0 brd 0.0.0.0
3: sit0@NONE: <NOARP> mtu 1480 qdisc noop state DOWN group default qlen 1000
    link/sit 0.0.0.0 brd 0.0.0.0
22: eth0@if23: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
~~~

~~~
PS C:\Users\alan> ping 172.17.0.2

正在 Ping 172.17.0.2 具有 32 字节的数据:
来自 172.17.0.2 的回复: 字节=32 时间=211ms TTL=59
来自 172.17.0.2 的回复: 字节=32 时间=75ms TTL=59
// 宿主机可以ping通
~~~



原理

每启动一个docker容器，docker就会给容器分配一个IP，我们只要安装了docker，就会有一个网卡桥接模式，使用的技术是veth-pair技术！

容器带来的网卡，都是一对一对的，一端连接协议，一段彼此相连。

~~~
以太网适配器 vEthernet (WSL):

   连接特定的 DNS 后缀 . . . . . . . :
   本地链接 IPv6 地址. . . . . . . . : fe80::39b3:a46a:d63a:a9ec%50
   IPv4 地址 . . . . . . . . . . . . : 172.28.128.1
   子网掩码  . . . . . . . . . . . . : 255.255.240.0
   默认网关. . . . . . . . . . . . . :
~~~

![image-20211013165607336](C:\Users\alan\AppData\Roaming\Typora\typora-user-images\image-20211013165607336.png)

重启会导致IP切换

docker exec -it tomcat02 ping tomcat01

ping不同

## --link

~~~
docker run -d -P --name tomcat03 --link tomcat02 tomcat
~~~

使用这种方式tomcat03可以ping通tomcat02,但是tomcat02ping不通tomcat03

--link 原理 
在/etc/hosts中，指向另一个容器

~~~
172.18.0.3  tomcat02 容器ID
~~~

不推荐使用 --link

docker0的问题：不支持容器名访问！



~~~
docker network  inspect 1c993e15a4d4
[
    {
        "Name": "bridge",
        "Id": "1c993e15a4d437af5153fbccfecddf8f390caae5ed178550ac4f3e655cc69325",
        "Created": "2021-10-13T02:05:50.141004693Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "df852c642704be5c25125c836eee76969954820f091fd5d75c27d93e74a9d545": {
                "Name": "centos_v1",
                "EndpointID": "a70f175209f49023c876c7bdcdc1ad3249cf7fc16d7ae7b53259e09471a118af",
                "MacAddress": "02:42:ac:11:00:02",
                "IPv4Address": "172.17.0.2/16",
                "IPv6Address": ""
            }
        },
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
~~~



## 自定义网络

//  查看所有网络

docker  network ls  



//   删除网络

docker network rm 网路名



### 网络模式

- bridge： 桥接（默认）
- none：    不配置网路
- host：     主机模式（和宿主机共享网络）
- container：  容器网络连通（用的少，局限大）
- 



默认带个   --net bridge   域名不能访问

~~~
docker run -d -P --name tomcat01 tomcat
docker run -d -P --name tomcat01  --net bridge tomcat
~~~



创建网络

~~~
docker network create  --driver  bridge  --subnet 192.168.0.0/16  --gateway 192.168.0.1  mynet
~~~

docker network inspect  网络ID 

~~~
[
    {
        "Name": "mynet",
        "Id": "1482cdae35c030f451e5173427301464039978846e8ca95c034180dfd3d37c2b",
        "Created": "2021-10-13T09:24:36.208004321Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.0.0/16",
                    "Gateway": "192.168.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {},
        "Labels": {}
    }
]
~~~



## 网路连通



docker network connect  [OPTIONS]  网络 容器名

连通之后， 将容器加入到该网络

一个容器两个IP

# DockerFile

~~~dockerfile
FROM centos // centos镜像
VOLUME ["/data1","/data2"]
CMD echo "finish, SUCCESS"
CMD /bin/bash
~~~

docker build  -f /xxx/dockerfile  -t alan/centos  .

生成alan/centos镜像



# 容器编排

定义，运行多个容器



docker-compose up
