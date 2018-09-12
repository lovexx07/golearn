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
 第二节
  import (
  	"flag"  接收参数的包
  	"fmt"
  )
  
  var a int
  
  func init() {
  	flag.IntVar(&a, "a", 123, "输入数字")
  }
  
  func main() {
  	flag.Parse()    解析命令参数，必须有这步
  	fmt.Println("hello world ", a)
  }
  go 语言的 输入命令
  
  第三节
  同一包下面的文件应该声明为 同一个package 名字，不然无法编译通过
  调用时可以go run demo4.go demo4_lib.go -name=jack 来调用，这里的方法可以是小写
  
  鸭子方法
  只要一个数据类型的方法集合中有接口Pet的所有方法（值方法与指针方法），那么它就一定是Pet接口的实现类型。这是一种无侵入式的接口实现方式
  
  接口变量赋值
  当我们给一个接口变量赋值的时候，该变量的动态类型毁于它的动态值一起被存储在一个专用的数据结构中。这个专用的数据结构叫做iface,iface的实例会包含2个指针，一个是指向类型的信息的指针，另一个是指向动态值的指针。

  