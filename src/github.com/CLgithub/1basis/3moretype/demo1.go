package main

import "fmt"

func main() {
	// pointer1()
	// struct1()
	// fmt.Println(v1, v2, v3, p, p4)
	array1()

}

//数组
func array1() {
	var a [10]int	//数组a是int类型的数组，长度为10
	//数组的长度是其类型的一部分，因此数组不能改变大小。 这看起来是一个制约，但是请不要担心； Go 提供了更加便利的方式来使用数组。
	fmt.Println(a)

	var b [2]string
	b[0]="Hello"
	b[1]="world"
	fmt.Println(b[0],b[1])
	fmt.Println(b)
	
}

//结构体文法
var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
	v4 = 1
	p4 = &v4
)

//一个结构体（ struct ）就是一个字段的集合
func struct1() {
	v := Vertex{1, 2}
	v.X = 4 //可以通过.来访问结构体
	p := &v //结构体指针，指向结构体的指针
	p.Y = 1e9
	fmt.Println(v)
}

//定义一个结构体
type Vertex struct {
	X int
	Y int
}

//指针
func pointer1() {
	i, j := 42, 2701
	p := &i         //&把i的指针赋给p，指针p存储的是i的地址
	fmt.Println(*p) //*操作符得到指针p指向的地址所存储的值
	fmt.Println(p)  //直接打印出p的值，即i的地址
	fmt.Println("---------------------")
	*p = 21         //将指针p指向的地址说存储的值设置为21
	fmt.Println(*p) //*操作符得到指针p指向的地址所存储的值
	fmt.Println(p)  //直接打印出p的值，即i的地址
	fmt.Println("----------------------")
	p = &j
	*p = *p / 37
	fmt.Println(j)
	fmt.Println(p)
}
