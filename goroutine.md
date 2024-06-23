## 什么是goroutine
goroutine是一种函数或方法，它与其他函数或方法同时运行。 goroutine可以算是轻量级的线程。 与线程相比,创建goroutine的成本很小。 因此，Go应用程序同时运行数千个goroutine是很常见的。  
goroutine使用go关键字创建，后根需要并发执行的函数调用。 当一个goroutine启动时，它与程序中的其他goroutine并发运行。 goroutine使用通道相互通信，通道是Go语言中的一种数据结构，允许在goroutine之间进行安全和同步的通信。 

## goroutine的优势
goroutine相对线程的优势如下。
* 与线程相比，goroutine对资源的消耗更少。
* goroutine被多路复用到较少数量的线程。 一个有数千个goroutine的程序中可能只有一个线程。
* goroutine利用通道通信。在使用goroutine访问共享内存时，通道设计可以防止竞争条件的出现。

## 如何启动一个goroutine
在Go语言中，每一个并发执行的活动称为goroutine，使用go关键字即可创建goroutine，形式如下
```
go funcName()
```
funcName()是我们定义好的函数或者闭包，将go剪剪子声明到一个需要调用的函数之前，在相同地址空间调用这个函数，这样该函数在执行时便会作为一个独立的并发goroutine。
```
package main
import "fmt"
func funcName(){
  fmt.Println("Hi goroutine")
}
func main(){
  go funcName()
  fmt.Println("main function")
}
```

## 如何启动多个goroutine
启动多个goroutine的程序如下
```
package main
import (
  "fmt"
  "time"
)
func printNumbers(){
  for i :=6 ; i<=8; i++ {
    time.Sleep(200 * time.Millisecond)
    time.Println("%c ", i)
  }
}
func printChars(){
  for i :='x'; i<='z'; i++ {
    time.Sleep(300 * time.Millisecond)
    time.Println("%c ", i)
  }
}
func main(){
  go printNumbers()
  go printChars()
  time.Sleep(2000 * time.Millisecond)
  time.Println("main end ")
}

```

## 通道 channel
使用chan关键字和数据类型类声明一个新的通道类型
```
var ch chan int
```
在以上代码中，ch是发送类型chan int的通道。 int通道的默认值是nil， 因此需要再使用之前对其进行赋值。 也可以使用make() 函数来声明和初始化通道。
```
ch := make(chan int)
```

## 发送和接收数据
使用通道发送和接收
```
package main
import "fmt"
func main(){
  n := 6
  //创建通道
  out := make(chan int)
  //启动goroutine
  go multiplyByTwo(n, out)
  //在此通道上收到任何输出后，将其打印到控制台并继续
  fmt.Println(<-out)
}
//这个函数现在接收一个通道作为它的第2个参数...
func multiplyByTwo(num int, out chan<- int) {
  result := num * 2
  // ... 并将结果通过管道传递给它
  out <- result
}
```

## 定向通道
通道可以是定向的。
```
out chan<- int
```
chan<- 声明，只能将数据发送到通道中，而不能从通道中接收数据
```
out <-chan int
```
<-chan 仅接收

## 阻塞条件
从通道发送或接收至的语句在它们自己的goroutine中阻塞
* 从通道接收数据的语句中阻塞，直到接收到一些数据
* 从通道发送数据的语句将等待发送的数据被接收

## select 语句
当有多个通道等待接收信息，并希望在其中任何一个通道并完成一个动作时，可以使用select语句。switch语句会从上到下遍历所有的case语句，并尝试找到与switch表达式的第一个case表达式。 一个找到匹配的case语句，它就退出并且不考虑其他case语句。 

## 缓冲通道
缓冲通道是一种内部具有存储容量的通道。 要常见缓存通道，可以向make()函数中添加第2个参数以指定容量：
```
out := make(chan int, 2)
```
out是一个容量为2的整数变量的缓冲通道。这意味着它在阻塞之前最多可以接收2个值。
