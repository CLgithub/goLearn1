package main

import "fmt"

func main() {
	// pointer1()
	// struct1()
	// fmt.Println(v1, v2, v3, p, p4)
	// array1()
	// slece1()
	// slice2()
	// slice3()
	slice4()

}

//构造slice make([]type,len,cap)
func slice4() {
	a:=make([]int,5)	//创建一个长度和容量都是5的数组，并返回一个slice指向这个数组
	printSlice("a",a)
	b:=make([]int,1,5)	//创建一个长度为1容量为5的数组，赋给b
	printSlice("b",b)
	c:=b[0:2]			//截取b的0到2位给c
	printSlice("c",c)
	d:=c[2:3]			//截取c的2到5位给d
	printSlice("d",d)
}

func printSlice(s string,x []int) {
	fmt.Printf("%s=%d len=%d cap=%d\n",s,x,len(x),cap(x))	
}

//slice2切片
//slice可以从新切片创建一个新的 slice 值指向相同的数组
// s[i:j]  表示slice s的第i个元素到第j个元素组成的新slice，包头不包尾
func slice3() {
	s1:=[]int{2,5,1,56,322,5}
	// fmt.Println(s1[1:7])
	fmt.Println(s1[1:6])
	fmt.Println(s1[1:3])
	fmt.Println(s1[1:2])
	fmt.Println(s1[1:1])
	// fmt.Println(s1[1:0])	 //	
}

//slice 的 slice (一个序列slice，其元素也是一个序列slice)
func slice2() {
	game := [][] string{
		[]string{"*","-","*"},
		[]string{"-","*","-"},
		[]string{"*","-","*"},
	}

	game[2][1]="#"

	for i:=0; i<len(game); i++{
		for j:=0; j<len(game[i]); j++{
			fmt.Print(game[i][j])
		}
		fmt.Println()
	} 	
}

//slice	片
/*
	一个 slice 会指向一个序列的值，并且包含了长度信息
	数组[]T 是一个元素类型为 T 的 slice，（数组是一个特殊的slice）
	len(s) 返回slice的长度
*/
func slece1() {
	a:=[]int{2,4,1,5}	
	fmt.Println("a ==",a)
	for i:=0; i<len(a); i++ {
		fmt.Printf("s[%d] == %d\n",i,a[i])
	}
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
