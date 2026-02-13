# dBackup: 华为云备份工具

一个基于华为云的备份工具，通过命令行的方式备份虚拟机、云硬盘、对象存储桶、快照等，当发生病毒入侵、人为误删除、软硬件故障等事件时，可通过备份将云资源数据恢复到任意还原点。

## 构建项目：

- 编译项目生成二进制文件：

```
[root@localhost ~]# git clone https://github.com/DongJunQAQ/dBackup.git
[root@localhost ~]# cd dBackup/
[root@localhost dBackup]# make
Start compiling this project on the  platform...
go build -o ./bin/dBackup ./main.go
Compilation completed: ./bin/dBackup
```

执行完此命令后会在./bin目录下生成名为dBackup的可执行文件；

- 验证是否安装成功：

```
[root@localhost dBackup]# cd ./bin/ && ./dBackup -v

      _ ____             _                
     | |  _ \           | |               
   __| | |_) | __ _  ___| | ___   _ _ __  
  / _' |  _ < / _` |/ __| |/ / | | | '_ \
 | (_| | |_) | (_| | (__|   <| |_| | |_) |
  \__,_|____/ \__,_|\___|_|\_\\__,_| .__/ 
                                   | |    
                                   |_|    

Version: v0.1.0
```

- 清除编译产物：

```
[root@localhost dBackup]# make clean
Cleaning compilation artifacts...
rm -rf ./bin
Cleanup completed
```



## 命令行自动补全：

```
[root@localhost bin]# ./dBackup completion bash > /etc/bash_completion.d/dBackup
```

执行完此命令后重新进入bash即可获取命令行自动补全功能；

如果是直接从Release页面下载的可执行文件则需要先将dBackup-linux-amd64重命名为dBackup之后再执行以上步骤。



## 快速开始：





## 使用指南：

### 获取帮助：

要获取一些帮助并了解dBackup的工作原理，您可以使用-h标志：

```
[root@localhost ~]# dBackup -h
一个基于华为云的备份工具，通过命令行的方式备份虚拟机、云硬盘、对象存储、快照等

Usage:
  dBackup [command]

Available Commands:
  backup      开始备份
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  login       登录华为云
  logout      登出华为云
  resource    查看相关资源
  vault       备份存储库相关操作

Flags:
  -h, --help      help for dBackup
  -v, --version   version for dBackup

Use "dBackup [command] --help" for more information about a command.
```

获取某个子命令的帮助信息：dBackup <subcommand> -h

```
[root@localhost ~]# dBackup vault -h
备份存储库相关操作，如创建、列出、添加资源等

Usage:
  dBackup vault [command]

Aliases:
  vault, v

Available Commands:
  add         添加资源至存储库
  create      创建备份存储库
  list        列出所有备份存储库

Flags:
  -h, --help   help for vault

Use "dBackup vault [command] --help" for more information about a command.
```



### 登录华为云：

使用AK（Access Key ID）和SK（Secret Access Key）登录至华为云并使用AES-GCM算法加密后存储至配置文件中。其是用户在华为云的长期身份凭证，华为云通过AK识别访问用户的身份，通过SK对请求数据进行签名验证，用于确保请求的机密性、完整性和请求者身份的正确性。通常来说AK和SK非常重要绝对不能发生泄露，因此dBackup使用动态生成的真随机密钥以及AES-GCM算法来加密AK与SK。

```
[root@localhost ~]# dBackup login 
请输入AK:                     
请输入SK:                                         
SUCCESS  AK/SK已安全保存至: /root/.dbackup_config.json
```



### 查询云资源ID：

在执行备份前通常需要使用此命令来查询要备份云资源的唯一ID。

```

```



### 存储库：

云备份使用存储库来存放备份。创建备份前，需要先创建至少一个存储库，并将服务器或磁盘绑定至存储库。服务器或磁盘产生的备份则会存放至绑定的存储库中。

- 创建存储库：

```

```

- 列出存储库：

```

```

- 将指定的云资源添加至存储库中：

```

```



### 执行备份：

对存储库执行备份，即生成备份还原点。

```

```



### 登出：

登出华为云并删除存放验证信息的配置文件。

```

```

