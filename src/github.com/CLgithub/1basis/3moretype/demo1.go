package main

import "fmt"


func main() {
	pointer1()
}

//指针
func pointer1() {
	i,j := 42,2701
	p:=&i 			//&把i的指针赋给p，指针p存储的是i的地址
	fmt.Println(*p)	//*操作符得到指针p指向的地址所存储的值
	fmt.Println(p)	//直接打印出p的值，即i的地址
	fmt.Println("---------------------")
	*p=21			//将指针p指向的地址说存储的值设置为21
	fmt.Println(*p)	//*操作符得到指针p指向的地址所存储的值
	fmt.Println(p)	//直接打印出p的值，即i的地址
	fmt.Println("----------------------")
	p=&j
	*p=*p/37
	fmt.Println(j)
	fmt.Println(p)	
}
