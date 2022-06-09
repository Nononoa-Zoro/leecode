### 包的引入规范

1. import的包按照 标准库 内部包 第三方包 由上到下依次排列
2. 不要使用相对路径引入包
3. 包名和Git路径不一致时，使用别名进行代替

### iota
iota是一个特殊的常量，可以被认为是编译器可以修改的常量。iota在关键字const出现时被重置为0，const中每新增一行常量声明，iota的值都会+1

### error处理

1. error作为返回值返回时，必须在返回值列表的最后一个
2. error返回应该独立判断，不要与其他变量组合判断



### panic的处理

1. 业务逻辑中禁用panic
2. main包里面只有程序完全不可运行才使用



### golang format

%v the value in a default format
when printing structs, the plus flag (%+v) adds field names
%#v a Go-syntax representation of the value
%T a Go-syntax representation of the type of the value
bool: %t
int, int8 etc.: %d
uint, uint8 etc.: %d, %#x if printed with %#v
float32, complex64, etc: %g
string: %s
chan: %p
pointer: %p



#### golang函数中的可变参数

```go
func a(i int,s ...string) {} // 可变长参数类型一定在最后一个函数入参
```



#### golang的闭包

捕获自由变量的绑定对函数文本执行“闭合”动作

接口只接受接口的嵌入



### channel

close(channel) 之后，可以读但无法写



### Golang变量的初始化

1. 包级别的变量初始化规则：确认依赖关系，优先初始化不依赖其他的变量
2. 函数内部的局部变量初始化规则：从上到下，从左到右，依次初始化。



### Golang程序初始化流程

1. 如果import了别的包，首先初始化导入的包。（直到某个包不依赖别的包）
2. 如果找到了一个包，它不依赖任何别的包，那么针对这个包下的go文件从上到下，依次
3. 初始化const
4. 初始化var
5. 执行init，对于同一个go文件，如果有多个init函数，执行顺序是从上到下。同包下多个go文件中的init函数执行顺序，根据文件名字排序，依次执行。

 const --> var --> init

<b>main函数执行之前的所有初始化操作都是在一个routine中执行的</b>



### 指针类型的方法接受者

如果p是一个Point类型的变量，但方法需要一个*Point接受者，可以简写为：

```go
p.ScaleBy(2) // 编译器会将p隐式转换为&p	
```



### 结构体内嵌类型

内嵌的字段会告诉编译器生成额外的包装方法来调用内嵌字段的方法

```go
func (p ColorPoint)Distance(q Point) float64 {
  return p.Point.Distance(q)
}
```



### 封装的优点

使用方不能直接修改对象变量，所以不需要更多语句来检查变量的值

隐藏实现细节可以防止被依赖属性改变导致使用方修改



### defer

defer可以在return和panic之前执行

可以定义多个defer函数，执行顺序是倒序

for循环里不要用defer defer只会在函数退出时候执行 代码块退出的时候不会执行



### Golang闭包惰性求值

Golang中的闭包也会保存外部环境，在闭包执行阶段会寻找外部环境最新的值处理



### map初始化

```go
func initMap(){
  var m map[int]int // 只是声明 nil
  m:=map[int]int{}  // 声明并初始化
  m:=make(map[int]int)
}
```



### slice和arr的转换

在go1.17之前

arr可以转换为slice 但是 slice不能转换为arr
go1.17对slice转换为arr做了支持



### for range 读 channel的值

```go
for c:=range ch {
  ....
}

如果不close(ch),会一直阻塞在这里，即便ch中没有数据也会阻塞
告诉其他消费者，以后不会向channel写数据了，消费完数据之后routine自动关闭
```



### select

select一直等待，直到一次通信来告知有一些情况可以执行，然后它执行这次通信。

多个case语句条件同时被满足，select会随机选择一个。

channel的零值是nil，在nil的channel上发送和接收都会阻塞。



### channel的关闭

关闭一个已经关闭的channel会产生Panic。优雅关闭channel会用到sync.Once

向一个已经关闭的channel发送数据会产生Panic



### channel操作表

| 操作  | nil的channel | 正常channel | 已关闭channel |
| :---: | :----------: | :---------: | :-----------: |
| <-ch  |     阻塞     | 成功或阻塞  |   读到零值    |
| ch<-  |     阻塞     | 成功或阻塞  |     panic     |
| close |    panic     |    成功     |     panic     |

### unsafe.Pointer

Unsafe.Pointer的特别之处在于，它可以绕过 Go 语言类型系统的检查，与任意的指针类型互相转换。也就是说，如果两种类型具有相同的内存结构（layout），我们可以将`unsafe.Pointer`当做桥梁，让这两种类型的指针相互转换，从而实现同一份内存拥有两种不同的**解读**方式。

### Golang的惰性求值

Golang中的闭包会保存外部环境，在闭包执行阶段会寻找外部环境最新的值处理。

### Golang Panic的情况

1. 空指针
2. 数组和切片访问越界
3. 通道操作 见channel操作表
4. 往未初始化的map写入数据
5. 类型断言失败，如果不用if ok的形式处理 会直接panic

#### Golang 内存模型

https://go.dev/ref/mem

https://www.modb.pro/db/88826

