package study

import (
	"errors"
	"fmt"
	"strconv"
)

func TestCast() {
	//字符串转数字, str不为纯数字时, 会发生syntax错误, go的error是作为返回值接收的, 使用_忽略
	str := "11"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("%.2f\n", float64(num))
	} else {
		fmt.Println(errors.Is(err, strconv.ErrSyntax))
	}
	//数字转字符串
	str = strconv.Itoa(num)
	println(str)

	//any / interface{}的强转
	var i any = "hello world"
	s, ok := i.(string)
	if ok {
		println(s)
	} else {
		println("转换失败")
	}
}
