package main

import "fmt"

type Duck interface {
	Quack()
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func main() {
	c := Cat{}
	c.Quack() // 可寻址结构体可以隐式调用指针方法 (&c).Quack()

	var d Duck = &Cat{} // 但结构体本身并不保护指针方法，因此这里必须要取地址才能转换为对应接口
	d.Quack()

}
