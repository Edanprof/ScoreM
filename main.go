package main

import(
	"fmt"
	"time"
	"strconv"
)

func main(){
	msg:="test"
	fmt.Println(msg)
	t:=time.Now()
	a:=strconv.Itoa(int(t.Month()))
	fmt.Println(a)
	fmt.Println("%d%d%d",t.Year(),int(t.Month()),t.Day())
}

