# 简单剖析`new`和`make`怎么用
基本认识:make函数为slice、map、chan这三种类型服务。日常开发中使用make初始化slice时要注意零值问题，否则又是一个p0事故

> 分配内存(new)

```go
// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.

func new(Type) *Type
````
大概翻译:新的内置函数分配内存。第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。

我们平常在使用指针的时候是需要分配内存空间的，未分配内存空间的指针直接使用会使程序崩溃，比如这样：
```go
var a *int64
*a = 10
````
我们声明了一个指针变量，直接就去使用它，就会使用程序触发panic，因为现在这个指针变量a在内存中没有块地址属于它，
就无法直接使用该指针变量，所以new函数的作用就出现了，通过new来分配一下内存，就没有问题了
```go
var a *int64 = new(int64)
 *a = 10
````

上面的例子，我们是针对普通类型int64进行new处理的，如果是复合类型，使用new会是什么样呢？来看一个示例：
```go
func main(){
 // 数组
 array := new([5]int64)
 fmt.Printf("array: %p %#v \n", &array, array)// array: 0xc0000ae018 &[5]int64{0, 0, 0, 0, 0}
 (*array)[0] = 1
 fmt.Printf("array: %p %#v \n", &array, array)// array: 0xc0000ae018 &[5]int64{1, 0, 0, 0, 0}
 
 // 切片
 slice := new([]int64)
 fmt.Printf("slice: %p %#v \n", &slice, slice) // slice: 0xc0000ae028 &[]int64(nil)
 (*slice)[0] = 1
 fmt.Printf("slice: %p %#v \n", &slice, slice) // panic: runtime error: index out of range [0] with length 0

 // map
 map1 := new(map[string]string)
 fmt.Printf("map1: %p %#v \n", &map1, map1) // map1: 0xc00000e038 &map[string]string(nil)
 (*map1)["key"] = "value"
 fmt.Printf("map1: %p %#v \n", &map1, map1) // panic: assignment to entry in nil map

 // channel
 channel := new(chan string)
 fmt.Printf("channel: %p %#v \n", &channel, channel) // channel: 0xc0000ae028 (*chan string)(0xc0000ae030) 
 channel <- "123" // Invalid operation: channel <- "123" (send to non-chan type *chan string) 
}
````
从运行结果可以看出，我们使用new函数分配内存后，只有数组在初始化后可以直接使用，slice、map、chan初始化后还是不能使用，会触发panic，
这是因为slice、map、chan基本数据结构是一个struct，也就是说他里面的成员变量仍未进行初始化，所以他们初始化要使用make来进行，make会初始化他们的内部结构，
我们下面一节细说。还是回到struct初始化的问题上，先看一个例子：
```go
type test struct {
 A *int64
}

func main(){
 t := new(test)
 *t.A = 10  // panic: runtime error: invalid memory address or nil pointer dereference
             // [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10a89fd]
 fmt.Println(t.A)
}
````
从运行结果得出使用new()函数初始化结构体时，我们只是初始化了struct这个类型的，而它的成员变量是没有初始化的，所以初始化结构体不建议使用new函数，
使用键值对进行初始化效果更佳。

`其实 new 函数在日常工程代码中是比较少见的，因为它是可以被代替，使用T{}方式更加便捷方便。`

> 初始化内置结构之make

make函数是专门支持 slice、map、channel 三种数据类型的内存创建，其官方定义如下
```go
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type: ..........

func make(t Type, size ...IntegerType) Type
````
大概翻译: make内置函数分配并初始化一个slice、map或chan类型的对象。像new函数一样，第一个参数是类型，而不是值。与new不同，make的返回类型与其参数的类型相同，
而不是指向它的指针。结果的取决于传入的类型

使用make初始化传入的类型也是不同的，具体可以这样区分：
```go
Func             Type T     res
make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
````
不同的类型初始化可以使用不同的姿势，主要区别主要是长度（len）和容量（cap）的指定，有的类型是没有容量这一说法，因此自然也就无法指定。
如果确定长度和容量大小，能很好节省内存空间。

demo
```go
func main(){
 slice := make([]int64, 3, 5)
 fmt.Println(slice) // [0 0 0]
 map1 := make(map[int64]bool, 5)
 fmt.Println(map1) // map[]
 channel := make(chan int, 1)
 fmt.Println(channel) // 0xc000066070
}
````
需要注意的点，就是`slice在进行初始化时，默认会给零值`，在开发中要注意这个问题，避免导致数据不一致

> new和make区别总结
- new函数主要是为类型申请一片内存空间，返回执行内存的指针
- make函数能够分配并初始化类型所需的内存空间和结构，返回复合类型的本身。
- make函数仅支持 channel、map、slice 三种类型，其他类型不可以使用使用make。
- new函数在日常开发中使用是比较少的，可以被替代。
- make函数初始化slice会初始化零值，日常开发要注意这个问题
> make函数底层实现

执行汇编指令：go tool compile -N -l -S file.go，我们可以看到make函数初始化slice、map、chan分别调用的是runtime.makeslice、runtime.makemap_small、runtime.makechan这三个方法，
因为不同类型底层数据结构不同，所以初始化方式也不同，我们只看一下slice的内部实现就好了
```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
 mem, overflow := math.MulUintptr(et.size, uintptr(cap))
 if overflow || mem > maxAlloc || len < 0 || len > cap {
  // NOTE: Produce a 'len out of range' error instead of a
  // 'cap out of range' error when someone does make([]T, bignumber).
  // 'cap out of range' is true too, but since the cap is only being
  // supplied implicitly, saying len is clearer.
  // See golang.org/issue/4085.
  mem, overflow := math.MulUintptr(et.size, uintptr(len))
  if overflow || mem > maxAlloc || len < 0 {
   panicmakeslicelen()
  }
  panicmakeslicecap()
 }

 return mallocgc(mem, et, true)
}
````
这个函数功能其实也比较简单(简单理解)：
- 检查切片占用的内存空间是否溢出。
- 调用mallocgc在堆上申请一片连续的内存。

检查内存空间这里是根据切片容量进行计算的，根据当前切片元素的大小与切片容量的乘积得出当前内存空间的大小，检查溢出的条件有四个：
- 内存空间大小溢出了
- 申请的内存空间大于最大可分配的内存
- 传入的len小于0，cap的大小只小于len

>new函数底层实现

new函数底层主要是调用runtime.newobject：
```go
// implementation of new builtin
// compiler (both frontend and SSA backend) knows the signature
// of this function
func newobject(typ *_type) unsafe.Pointer {
 return mallocgc(typ.size, typ, true)
}
````
内部实现就是直接调用mallocgc函数去堆上申请内存，返回值是指针类型。