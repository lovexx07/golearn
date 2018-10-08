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

  panic 意思是抛出一个异常， 和python的raise用法类似
  recover是捕获异常，和python的except用法类似
  defer会延迟函数到其他函数之后完之后再执行，后面跟的是函数
  golang 的错误处理流程：当一个函数在执行过程中出现了异常或遇到
  panic()，正常语句就会立即终止，然后执行 defer 语句，再报告异
  常信息，最后退出 goroutine。如果在 defer 中使用了 recover() 
  函数,则会捕获错误信息，使该错误信息终止报告。
  
  多个defer 执行顺序，按照编译好的顺序从后往前执行
  
  
  go test golearn/test20/q1  进行功能测试
  
  go test -bench=. -run=^$ golearn/test20/q2 性能测试
    -bench=.  只有加了这个标记，命令才会进行性能测试。表示需要执行任意名称的性能测试函数  正则
    -run^$ 表示执行哪些功能测试函数，这里表示不执行任何功能测试函数  正则
    
    goos: darwin
    goarch: amd64
    pkg: golearn/test20/q2
    BenchmarkGetPrimes-4      300000              4277 ns/op
    # 可以同时运行goruntine逻辑cpu数     简称该值为执行次数       单次执行GetPrimes函数的平均耗时
    PASS
    ok      golearn/test20/q2       1.338s

    //                       count：在不同情况下执行次数  cpu：设置不同cpu（ 1， 2， 3， 4。。。）
    go test -bench=. -run=^$  -count=2  -cpu=2,4  golearn/test21/q0

  