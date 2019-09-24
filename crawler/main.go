package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"regexp"

	//"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	// 检测编码
	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	// 会包括头部的信息
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printCityList(all)



	fmt.Printf("%s\n", all)

}


func determineEncoding(r io.Reader) encoding.Encoding {
	// 这里如果 resp.Body 被读了后，就不能再读了，所以我们需要 bufio.NewReader()
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	// e encoding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+"[^>]*>[^<]+</a>`)
	match := re.FindAll(contents, -1)
	fmt.Println(match)

}