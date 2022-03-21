# FastFileSafe_保密文件加密工具

#### 介绍
为企业解决公网传递保密文件带来安全隐患问题

针对Windows版本特别提供了注册表右键注入方案。

#### 软件架构
Golang语言,使用AES+ZIP对文件进行加密压缩,支持跨平台linux,windows,macos等

go version go1.17.7

#### 使用说明
Windows 命令行模式:

1.双击FastFileSafe.exe运行

2.根据提示输入文件路径(可文件或目录)

3.程序返回加密密钥,没密钥将永远无法打开文件



Windows右键模式:
1.用文本工具打开右键注册.reg 

2.修改F:\\golang\\git\\FastFileSafe\\   路径为FastFileSafe.exe目录位置

3.双击"右键注册.reg"注册

4.在文件或文件夹上右键可看见文件加密/解密 选项(如果是dooxb结尾的文件自动识别成解密模式)

5.加密成功将会返回文件密钥,没密钥将永远无法打开文件


Linux或MacOS
1.给文件赋权限chmod 755 FastFileSafe

2.设置当前用户环境变量  ln -s /路径/FastFileSafe /usr/local/bin/

3.  ./FastFileSafe /opt/xxx.jpg  进行文件加解密 (如果是dooxb结尾的文件自动识别成解密模式)


###普通用户使用说明
1.Windows环境右键加密文件
![右键加密文件](https://gitee.com/aliu/FastFileSafe/raw/develop/jpg/1.%E6%99%AE%E9%80%9A%E6%96%87%E4%BB%B6%E5%8A%A0%E5%AF%86-1.png "右键加密文件")

2.在<font color=red>**dooxb结尾**</font>文件右键解密文件,输入加密时的密码

![右键加密文件](https://gitee.com/aliu/FastFileSafe/raw/develop/jpg/1.%E6%99%AE%E9%80%9A%E6%96%87%E4%BB%B6%E5%8A%A0%E5%AF%86-2.png "右键加密文件")
 

#### 源码编译说明
go mod vendor -v

SET CGO_ENABLED=0

SET GOARCH=amd64

SET GOOS=darwin

go build -o FastFileSafe-macos


SET GOOS=windows

go build -ldflags "-w"



SET GOOS=linux

go build -o FastFileSafe-linux




SET GOOS=linux

SET GOARCH=arm64

go build -o FastFileSafe-linux-arm64


或执行build.bat


###360等误报说明
不知为什么打包后360一会儿就开始误报,大家不放心可以自己看源码自己打包也行,反正我这打包都是从官网下SDK.