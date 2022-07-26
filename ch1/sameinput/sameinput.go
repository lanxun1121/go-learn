package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	// 屏幕输入
	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	if input.Text() == "0" {
	//		break
	//	}
	//	count[input.Text()]++
	//}
	// 文件读取
	if len(os.Args) < 2 {
		fmt.Println("ERROR input")
		return
	}
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		println(input.Text())
		count[input.Text()]++
	}
	for text, num := range count {
		if num > 1 {
			fmt.Println(text, num)
		}

	}
}
