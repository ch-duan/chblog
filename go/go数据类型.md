# Go数据类型

#### Boolean

#### 数值类型

整数、浮点数或复数类型分别表示 *整数* 、浮点数或复数值的集合。它们统称为数字类型。预先声明的独立于体系结构的数字类型是：

```
uint8       the set of all unsigned  8-bit integers (0 to 255) (0 to 2^8-1)
uint16      the set of all unsigned 16-bit integers (0 to 65535) (0 to 2^16-1)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295) (0 to 2^32-1)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615) (0 to 2^64-1)

int8        the set of all signed  8-bit integers (-128 to 127) (-2^7 to 2^7-1)
int16       the set of all signed 16-bit integers (-32768 to 32767) (-2^15 to 2^15-1)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647) (-2^32 to 2^32-1)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807) (-2^63 to 2^63-1)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

n位整数 的值是n位宽，并使用 [二进制补码算法](https://en.wikipedia.org/wiki/Two's_complement)表示。

还有一组预先声明的具有特定于实现大小的整数类型：

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

#### String

字符串*类型*表示字符串值的集合。字符串值是一个（可能为空的）字节序列。字节数称为字符串的长度，永远不会是负数。字符串是不可变的：一旦创建，就不可能更改字符串的内容。预先声明的字符串类型是**string**; 它是一个[定义的类型](https://go.dev/ref/spec#Type_definitions)。

可以用内置方法**len**来获取长度，可以通过速度索引**0**到**len(s)-1**。不支持取元素地址操作

#### 数组类型

数组是单一类型的元素的编号序列，称为元素类型。元素的数量称为数组的长度，并且永远不会是负数。可以使用内置方法**len**来计算长度

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))

var s [32]int //声明一个长度为32的int数组
var s = [32]int{1,2,3,4} //声明一个长度为32的int数组并初始化部分数据
s := [32]int{1,2,3,4}  //声明一个长度为32的int数组并初始化部分数据
```

#### slice

切片是*底层数组*的连续段的描述符，并提供对该数组中元素的编号序列的访问。切片类型表示其元素类型的所有数组切片的集合。元素的数量称为切片的长度，并且永远不会是负数。未初始化切片的值为 `nil`。

#### map

默认值为nil

#### channel

默认值为nil

#### 自定义类型

关键字 `type`基于现有类型定义用户自定义类型。

```
type A int //定义新类型
type B = int //别名（实际类型还是int）
```

* 即使底层类型相同，也非同一类型（区别于别名）
* 除运算符外，不继承任何信息（方法等）
* 不能隐式转换，不能直接比较。

#### 未命名类型

与 `bool`、`int` 这些有明确标识的类型不同，`array` 、`slice` 、`map` 、`channel` 等与其元素或长度属性相关被称为未命名类型。

```
[] int //unamed type
type A [] int //named type
[]int, [] byte //元素类型不同，不是同一类型

func main(){
  var a,b interface{
    test()
  }
  fmt.Println(a==b) // true
}

```

###### 具有**相同声明的未命名类型被视作同一类型：**

* 相同基类型的指针。
* 相同元素类型和长度的数组。
* 相同元素类型的切片。
* 相同键值类型的字典(map)。
* 相同数据类型及操作方向的通道(channel)。
* 相同字段序列（字段名、类型、标签、顺序）的结构体。
* 相同签名（参数和返回值 列表，不包括参数名）的函数（function）。
* 相同方法集（方法名、方法签名，不包括顺序）的接口（interface）。
