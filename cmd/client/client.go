package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/mukuro690528/chatroom/pkg/user"
)

func main() {
	conn, err := rpc.Dial("tcp", "0.0.0.0:5288")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	if err := conn.Call("HelloService.Hello", "Sherry", &reply); err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	var info *user.Info
	if err := conn.Call("HelloService.UserUnfo", "001", &info); err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)

}
