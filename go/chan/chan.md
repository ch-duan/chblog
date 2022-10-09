## channel

通道提供了一种机制，用于并发执行函数，以便通过发送和接收指定元素类型的值进行通信。未初始化通道的值为 。`nil`

可选运算符指定通道方向、发送或接收。如果给出了方向，则信道是定向的，否则它是双向的。通道可能被限制为仅通过赋值或显式转换发送或接收。`<-`

```
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

运算符与最左边的可能关联：`<-chan`

```
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

可以使用内置函数 [`make`](https://go.dev/ref/spec#Making_slices_maps_and_channels) 创建一个新的、初始化的通道值，该函数采用通道类型和可选容量作为参数：

```
make(chan int, 100)
```

容量（以元素数表示）设置通道中缓冲区的大小。如果容量为零或不存在，则信道无缓冲，仅当发送方和接收方都准备就绪时，通信才会成功。否则，如果缓冲区未满（发送）或不为空（接收），则通道将被缓冲并且通信成功而不阻塞。通道永远不会准备好进行通信。`nil`

通道可以使用内置功能[关闭来关闭](https://go.dev/ref/spec#Close)。[接收运算符](https://go.dev/ref/spec#Receive_operator)的多值赋值窗体报告在通道关闭之前是否发送了接收值。

单个通道可用于[发送语句](https://go.dev/ref/spec#Send_statements)、[接收操作](https://go.dev/ref/spec#Receive_operator)以及通过任意数量的 goroutine 调用内置函数 [`cap`](https://go.dev/ref/spec#Length_and_capacity) 和 [`len`](https://go.dev/ref/spec#Length_and_capacity)，而无需进一步同步。通道充当先进先出队列。例如，如果一个goroutine在通道上发送值，而另一个**goroutine**接收这些值，则值将按发送的顺序接收。
