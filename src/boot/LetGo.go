package main

import (
	"fmt"
	"net/http"
	"../util"
	"reflect"
)

func letGo() {
	res, _ := http.Get("https://studygolang.com/articles/9467")
	fmt.Println(res.Status)
	fmt.Println(res.Header.Get("Content-Type"))
}

func println(obj ...interface{}) {
	fmt.Println(obj)
}

var str string = "hello world"

func testCode1() {
	//fmt.Println("Hello, 世界")
	//for i := 0; i < 5; i++ {
	//	go letGo()
	//}

	str2 := "66666"

	println(len(str2))
	println(len(str))
	println(str)

	var count int = len(str)

	// i:=0 声明并赋值
	for i := 0; i < count; i++ {
		println(string(str[i]))
	}

	//声明
	var j int
	//赋值
	for j = 0; j < count; j++ {
		println(str[j])
	}

	//定义常量
	const pi float32 = 3.14
	//常量无法再次赋值
	//pi = 66.6

	println(pi)

	const (
		c0 = iota
		c1 = iota
		c2 = iota
		c3 = iota
	)

	println(c0)
	println(c1)
	println(c2)
	println(c3)

	//常量可以自己推到
	const (
		c01 = iota
		c11
		c21
		c31
		c41 = 8
		c51 //赋值为上面数值8
	)
	println(c01)
	println(c11)
	println(c21)
	println(c31)
	println(c41)
	println(c51)
	println(string(99))

	var nums = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}

	println(util.Sum(5, 6))



	//fmt.Println(time.Now())
	//time.Sleep(10 * time.Second)
	//fmt.Println(time.Now())
}
// 基于现有数组创建分片
func createSlice1() {
	var nums [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(nums)
	println(reflect.TypeOf(nums).String())

	slice1 := nums[:5]
	var slice2 []int = nums[:5]

	println(slice1)
	println(slice2)
	println(reflect.TypeOf(slice1).String())
	println(reflect.TypeOf(slice2).String())

}

//[size]int是数组
//[]int是分片
func createSlice2() {
	var slice []int = make([]int, 5)
	println(slice)

	//size=5，预留10个空间,这10个空间包含面前5个
	var slice2 []int = make([]int, 5, 10)
	println(slice2)

	//初始化创建
	var slice3 []int = []int{1, 2, 3, 4, 5}
	println(slice3)

	println("===================================")
	println(len(slice))
	println(len(slice2))
	println(len(slice3))
	println("===================================")
	println(cap(slice))
	println(cap(slice2))
	println(cap(slice3))
	println("===================================")

	// 分片的append返回一个新的分片
	slice4 := append(slice3, 6, 7, 8)

	println("slice3 len = ", len(slice3))
	println("slice3 cap", cap(slice3))
	println("slice4 len = ", len(slice4))
	println("slice4 cap", cap(slice4))

	println("===================================")
	delIndex := 5
	//append后面三个点类似于scala可变参数传入数组时候后面需要:_* 一样
	slice5 := append(slice4[:delIndex - 1], slice4[delIndex:]...)
	println(slice5)
}

func createMap() {
	var numMap map[int]int = make(map[int]int)
	numMap[1] = 1
	numMap[2] = 2
	numMap[3] = 3
	println(numMap)

	key := 1
	//返回值两个，一个key对应的值，另一个是是否存在
	get, existed := numMap[key]
	if existed {
		println("key = ", key, "value = ", get, "existed!")
	} else {
		//当key不存在时候，返回的value是该类型对应的一个默认值
		println("key = ", key, "not existed!")
	}

	//删除元素

	delete(numMap, 3)

	println(numMap)

	//map size
	println(len(numMap))

}
//switch control
func switchCtrl(i int, j int) {

	// go switch 每一个case 不需要break，会自动break
	//如果某一个case执行之后向继续随后的break的话可以使用fallthrough关键字
	switch i {
	case 0:println(0)
	case 1 :println(1)
	case 2 :fallthrough
	case 3: println(3)
	default:
		println("default case !")
	}

	// switch
	switch  {
	case j > 0 && j < 10:
		println("一位数！")
	case j > 99 && j < 100:
		println("两位数！")
	}

	println("===============")
}

func gotoCtrl2(ok bool) string {

	if ok {
		goto here
	} else {
		println("not goto")
		return "not goto here!"
	}
	here:{
		println("goto")
		return "go here"
	}
}

func gotoCtrl1(ok bool) string {

	// 标签代码，会执行。
	here:{
		println("goto")
		return "go here"
	}
	//如下代码不会执行，被here标签代码中断返回了。
	if ok {
		goto here
	} else {
		println("not goto")
		return "not goto here!"
	}

}

func channleCtrl() {

	// 创建一个大小为1的channel，言外之意就是没有缓冲
	// 这种channel只有读写同时都准备好了才可以进行操作，否则阻塞
	ch := make(chan int, 1)

	for {
		select {
		// IO要向channel发起写 0 操作，一旦IO可操作就会被select执行写
		case ch <- 0:
			println("chan <- 0")
		// the same to adove
		case ch <- 1:
			println("chan <- 1")
		}
		// read data from channel
		i := <-ch
		fmt.Println(i)
	}
}

//带有缓冲的Buffer
func channelBuffer() {
	bufferChannel := make(chan int, 256)
	for {
		select {
		// IO要向channel发起写 0 操作，一旦IO可操作就会被select执行写
		case bufferChannel <- 0:
			println("chan <- 0")
		// the same to adove
		case bufferChannel <- 1:
			println("chan <- 1")
		}
		println("channelBuffer size = ", len(bufferChannel))
		// read data from channel
		i := <-bufferChannel
		fmt.Println(i)
	}
}

type Vector []float64
// 分配给每一个CPU的计算任务
//是为Vetcot添加方法
func (v Vector) DoSome(i, n int, u Vector, c chan float64) {
	sum := 0.0
	for ; i < n; i++ {
		//v[i] += u.Op(v[i])//书籍上面写的OP没看懂
		sum += v[i]
	}
	println(sum)
	c <- sum
	// 发信号告诉任务管理者我已经计算完毕了
}

const NCPU = 2
// 如果总共同拥有16核
func (v Vector) DoAll(u Vector) float64 {
	c := make(chan float64, NCPU)  // 用于接收每一个CPU的任务完毕信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i * len(v) / NCPU, (i + 1) * len(v) / NCPU, u, c)
	}
	sum := 0.0
	// 等待全部CPU的任务完毕
	for i := 0; i < NCPU; i++ {
		sum += <-c    // 获取到一个数据，表示一个CPU计算完毕了
	}
	// 到这里表示全部计算已经结束
	return sum
}

func main() {

	vector := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	u := make(Vector, 10)
	sum := vector.DoAll(u)
	println(sum)

}