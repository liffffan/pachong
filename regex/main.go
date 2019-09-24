package main

import (
	"fmt"
	"regexp"
)

const text =
	`
My email is ccmouse@gmail.com dada
email is abc@dada.com
email2 is 			kkk@qq.com.cn

`



func main() {
	// re 正则表达式匹配器
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllString(text, -1)
	match2 := re.FindAllStringSubmatch(text, -1)
	var url map[string]string = make(map[string]string, 100)
	for _,m := range match2 {
		for _,v := range m {
			fmt.Println(v)
			url["url"] = v
		}
	}
	fmt.Println(match)
	fmt.Println(match2)
	fmt.Println(url)

}
