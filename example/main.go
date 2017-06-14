package main

import (
	"fmt"
	"github.com/xiaoxiaoyijian123456/go3n1"
	"time"
)

func main() {
	t1 := time.Now()
	i := 1
	j := 10000000
	k, maxlen := go3n1.Maxlen3n1(uint64(i), uint64(j))
	elapsed := time.Since(t1)
	fmt.Println("Elapsed: ", elapsed)
	fmt.Printf("num = %d, maxlen=%d\n", k, maxlen)
	/*
		list := go3n1.List3n1(k)
		fmt.Print("list: [")
		for v := range list {
			fmt.Printf("%d,", v)
		}
		fmt.Print("]")
	*/
}
