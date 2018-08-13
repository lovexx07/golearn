# golearn
go 语音自学
第一节 环境配置与hello world
  使用mac去官网下载安装包即可，安装完后,打开命令工具  输入 >>>  go version 有反应则安装成功
  基本命令 
    go build 编译代码
    go run test.go 运行
  test.go 文件基本格式为
    package main     入口包，固定格式，程序入口
    import "fmt"     导包（基本包）
    func main(){     入口方法
      fmt.Println("hello world")   输出语句
    }
  ---------------------------------
  运行 go build------>  go run test.go  ----》hello world
 
 

 第二节 
  Go 环境配置 
  mac 找到根目录pengxin  下面的 .bash_profile 文件
  使用 vim .bash_profile 打开，按i 进入输入模式 输入下面部分
  export GOPATH=/Users/pengxin/Documents/workspace/workspace_go 
  export GOBIN=GOPATH/bin
  exportPATH=PATH:$GOBIN 
  shift + : 结束输入状态，输入w，保存退出
  source ~/.bash_profile 使设置马上生效
  go env 就可以查看是否配置成功
  ******** 
