package main

import (
	"fmt"
	"math"
)

//测试for循环
func testFor1() {
	sum :=0
	for i := 0; i < 10; i++ {	//不需要小括号，大括号是必须的
		sum+=i
	}
	fmt.Println(sum)	
}
//测试for循环2
func testFor2() {
	// sum :=1
	var sum int=1
	// for ; sum < 1000; {	//初始化语句和后置语句是可选的
	// for sum<1000 {			//初始化语句和后置语句是可选的，把;也去掉就相当于while语句
	for {					//把中间的判断预计也去掉，就变成了while(true)
		sum+=sum	
	}
	fmt.Println(sum)
}

//测试if
func testIf1() {
	var a,b float64=2,-4
	fmt.Println(sqrt(a),sqrt(b))
}
func sqrt(i float64) string {
	if i<0{
		return sqrt(-i)+"i"
	}	
	return fmt.Sprint(math.Sqrt(i))		//Sprint，格式使用其操作数的默认格式，并返回结果字符串。当两者都不是字符串时，在操作数之间添加空格
	// fmt.Printf("%T %v \n",math.Sqrt(i),math.Sqrt(i))
	// return "xxx"
}
func testIf2() {
	fmt.Println(
		pow(3,2,10,
			),	
		pow(3,3,20),		//这种写多行的，最后一行都必须加;
	)	
}
func pow(a,b,c float64) float64 {
	if i:=math.Pow(a,b); i<c {	//math.Pow(a,b) 求a^b
		return i
	} else {
		fmt.Printf("%g >= %g\n",i,c)
	}
	//这里开始就不能使用i了
	return c
}
//练习
func Sqrt(x float64) {
	fmt.Printf("%g 的平方根是 %g \n", x,math.Sqrt(x))
	z0:=1.0
	// for z0=float64(z0-((z0*z0-x)/(2*z0)));math.Sqrt(x)!=z0; {
	// 	z0=float64(z0-((z0*z0-x)/(2*z0)))
	// 	fmt.Println(z0)
	// }

	// for i := 0; i < 10; i++ {
	// 	z0=z0-((z0*z0-x)/(2*z0))
	// 	fmt.Println(z0)
	// }

	
	for z1:=z0-((z0*z0-x)/(2*z0)) ; z0!=z1;  {	//运输完后把
		z0=z1 						//运算前的先暂存在z0
		z1=z1-((z1*z1-x)/(2*z1))	//运输后的为z1
		fmt.Println(z0)
		fmt.Println(z1)
		z1=z1 						//运输完后把运输后的赋给z1,准备下次运输
	}
}
func main() {
	// testFor2()	
	// fmt.Println(fmt.Sprint('a'),fmt.Sprint("a"))
	// testIf2()
	Sqrt(4)
}