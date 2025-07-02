package main

import (
	"fmt"
	"time"
)

// var x, y int
// var ( // 这种因式分解关键字的写法一般用于声明全局变量
//
//	a int
//	b bool
//
// )
//
// var c, d int = 1, 2
// var e, f = 123, "hello"
// 这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello"
// const (
//
//	i = 1 << iota
//	j = 3 << iota
//	k
//	l
//
// )
// 定义接口
//type Shape interface {
//	Area() float64
//	Perimeter() float64
//}
//
//// 定义一个结构体
//type Circle struct {
//	Radius float64
//}
//
//// Circle 实现 Shape 接口
//func (c Circle) Area() float64 {
//	return math.Pi * c.Radius * c.Radius
//}
//
//func (c Circle) Perimeter() float64 {
//	return 2 * math.Pi * c.Radius

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	//fmt.Println("Hello, World!")
	//var x int
	//const Pi float64 = 3.14159265358979323846
	//if x > 0 {
	//
	//}
	//var stockcode = 123
	//var enddate = "2020-12-31"
	//var url = "Code=%d&endDate=%s"
	//var target_url = fmt.Sprintf(url, stockcode, enddate)
	//fmt.Println(target_url)
	//intVal := 1
	//fmt.Println(intVal)
	//g, h := 123, "hello"
	//fmt.Println(x, y, a, b, c, d, e, f, g, h)
	//var a string = "abc"
	//fmt.Printf("hello, %s world", a)
	//const (
	//	a = iota //0
	//	b        //1
	//	c        //2
	//	d = "ha" //独立值，iota += 1
	//	e        //"ha"   iota += 1
	//	f = 100  //iota +=1
	//	g        //100  iota +=1
	//	h = iota //7,恢复计数
	//	i        //8
	//)
	//fmt.Println(a, b, c, d, e, f, g, h, i)
	//fmt.Println("i=", i)
	//fmt.Println("j=", j)
	//fmt.Println("k=", k)
	//fmt.Println("l=", l)
	//var a int = 4
	//var b int32
	//var c float32
	//var ptr *int
	//
	///* 运算符实例 */
	//fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	//fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	//fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)
	//
	///*  & 和 * 运算符实例 */
	//ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	//fmt.Printf("a 的值为  %d\n", a)
	////fmt.Printf("*ptr 为 %d\n", *ptr)
	//var a int = 20
	//var b int = 10
	//var c int = 15
	//var d int = 5
	//var e int
	//
	//e = (a + b) * c / d // ( 30 * 15 ) / 5
	//fmt.Printf("(a + b) * c / d 的值为 : %d\n", e)
	//
	//e = ((a + b) * c) / d // (30 * 15 ) / 5
	//fmt.Printf("((a + b) * c) / d 的值为  : %d\n", e)
	//
	//e = (a + b) * (c / d) // (30) * (15/5)
	//fmt.Printf("(a + b) * (c / d) 的值为  : %d\n", e)
	//
	//e = a + (b*c)/d //  20 + (150/5)
	//fmt.Printf("a + (b * c) / d 的值为  : %d\n", e)
	// 定义两个通道
	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//
	//// 启动两个 goroutine，分别从两个通道中获取数据
	//go func() {
	//	for {
	//		ch1 <- "from 1"
	//	}
	//}()
	//go func() {
	//	for {
	//		ch2 <- "from 2"
	//	}
	//}()
	//
	//// 使用 select 语句非阻塞地从两个通道中获取数据
	//for {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println(msg1)
	//	case msg2 := <-ch2:
	//		fmt.Println(msg2)
	//	default:
	//		// 如果两个通道都没有可用的数据，则执行这里的语句
	//		fmt.Println("no message received")
	//	}
	//}
	//sum := 0
	//for i := 0; i <= 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)
	//sum := 1
	//for sum <= 10 {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//// 这样写也可以，更像 While 语句形式
	//for sum <= 10 {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//map1 := make(map[int]float32)
	//map1[1] = 1.0
	//map1[2] = 2.0
	//map1[3] = 3.0
	//map1[4] = 4.0
	// 读取 key 和 value
	//for key, value := range map1 {
	//	fmt.Printf("key is: %d - value is: %f\n", key, value)
	//}
	//
	//// 读取 key
	//for key := range map1 {
	//	fmt.Printf("key is: %d\n", key)
	//}
	//
	//// 读取 value
	//for _, value := range map1 {
	//	fmt.Printf("value is: %f\n", value)
	//}
	// 定义一个匿名函数并将其赋值给变量add
	// 将匿名函数作为参数传递给其他函数
	//calculate := func(operation func(int, int) int, x, y int) int {
	//	return operation(x, y)
	//}
	//// 也可以直接在函数调用中定义匿名函数
	//difference := calculate(func(a, b int) int {
	//	return a - b
	//}, 10, 4)
	//fmt.Println("10 - 4 =", difference)
	//balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	////  将索引为 1 和 3 的元素初始化
	//balance1 := [5]float32{1: 2.0, 3: 7.0}
	//fmt.Println(balance)
	//fmt.Println(balance1)
	//s := []int{1, 2, 3}
	//var numbers []int
	//printSlice(numbers)
	//
	///* 允许追加空切片 */
	//numbers = append(numbers, 0)
	//printSlice(numbers)
	//
	///* 向切片添加一个元素 */
	//numbers = append(numbers, 1)
	//printSlice(numbers)
	//
	///* 同时添加多个元素 */
	//numbers = append(numbers, 2, 3, 4)
	//printSlice(numbers)
	//
	///* 创建切片 numbers1 是之前切片的两倍容量*/
	//numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	//
	///* 拷贝 numbers 的内容到 numbers1 */
	//copy(numbers1, numbers)
	//printSlice(numbers1)
	// 使用字面量创建 Map
	//m := map[string]int{
	//	"apple":  1,
	//	"banana": 2,
	//	"orange": 3,
	//}
	//// 获取键值对
	//v1 := m["apple"]
	//v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值
	// 遍历 Map
	//for k, v := range m {
	//	fmt.Printf("key=%s, value=%d\n", k, v)
	//}
	//fmt.Println(v1, v2, ok)
	//num := 3.14
	//str := strconv.FormatFloat(num, 'f', 2, 64)
	//fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
	//c := Circle{Radius: 5}
	//var s Shape = c // 接口变量可以存储实现了接口的类型
	//fmt.Println("Area:", s.Area())
	//fmt.Println("Perimeter:", s.Perimeter())
	//var i interface{} = "42"
	//if str, ok := i.(string); ok {
	//	fmt.Println("String:", str, ok)
	//} else {
	//	fmt.Println("Not a string")
	//}
	//var phone Phone
	//
	//phone = new(NokiaPhone)
	//phone.call()
	//
	//phone = new(IPhone)
	//phone.call()
	//err := errors.New("this is an error")
	//fmt.Println(err) // 输出：this is an error
	//result, err := divide(10, 0)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Result:", result)
	//}
	//go sayHello() // 启动 Goroutine
	//for i := 0; i < 5; i++ {
	//	fmt.Println("Main")
	//	time.Sleep(100 * time.Millisecond)
	//}
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从通道 c 中接收
	//
	//fmt.Println(x, y, x+y)
	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	//// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	//// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	//// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	//for i := range c {
	//	fmt.Println(i)
	//}
	//var wg sync.WaitGroup
	//
	//for i := 1; i <= 3; i++ {
	//	wg.Add(1) // 增加计数器
	//	go worker(i, &wg)
	//}
	//
	//wg.Wait() // 等待所有 Goroutine 完成
	//fmt.Println("All workers done")
	m := make(map[int]int)

	//写goroutine
	go func() {
		for i := 0; i < 100000; i++ {
			m[i] = i // 写操作
		}
	}()

	// 读goroutine
	go func() {
		for i := 0; i < 1000000; i++ {
			_ = m[i] // 读操作
		}
	}()

	time.Sleep(time.Second)
}

//	func printSlice(x []int) {
//		fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
//	}
//
//	func divide(a, b int) (int, error) {
//		if b == 0 {
//			return 0, errors.New("division by zero")
//		}
//		return a / b, nil
//	}
//
//	func sayHello() {
//		for i := 0; i < 5; i++ {
//			fmt.Println("Hello")
//			time.Sleep(100 * time.Millisecond)
//
//		}
//	}
//
//	func sum(s []int, c chan int) {
//		sum := 0
//		for _, v := range s {
//			sum += v
//		}
//		//time.Sleep(100 * time.Millisecond)
//		fmt.Printf("Sending sum: %d from slice %v\n", sum, s)
//		c <- sum // 把 sum 发送到通道 c
//	}
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	close(c)
//}
//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done() // Goroutine 完成时调用 Done()
//	fmt.Printf("Worker %d started\n", id)
//	fmt.Printf("Worker %d finished\n", id)
//}
//
//type student struct {
//	name string
//}
