package main

import (
	"fmt"
	"net/url"
)

func main() {
	uri, _ := url.Parse("https://qq.com/http://www.shopee.tw")
	fmt.Println(uri.Host)
}
