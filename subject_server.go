package main

import (
	"github.com/qianlidongfeng/subject/server"
	"github.com/qianlidongfeng/log"
)
func main(){
	server := server.New()
	err:=server.Run()
	if err != nil{
		log.Fatal(err)
	}
}
