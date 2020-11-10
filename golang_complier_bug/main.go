// Ref: https://mp.weixin.qq.com/s/Tyl6dSb7mHBuqqN6WvEuaw
package main

import "fmt"

func main() {
	rates := []int32{1, 2, 3, 4, 5, 6}
	for star, rate := range rates {
		if star+1 < 1 { // 将这里的+1换成+2可避免这个问题
			panic("")
		}
		fmt.Println(star, rate)
	}

}
