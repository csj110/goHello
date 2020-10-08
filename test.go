package main

import (
	"fmt"
	"math/rand"
)

func test() {
	for i := 0;i<30;i++ {
		fmt.Println(genCode())
	}
}

func genCode() string {
	code :=rand.Intn(1000000)
	codeString:=fmt.Sprintf("%06d",code)
	return codeString
}