package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	con, err := net.Dial("tcp", "rickandmortyapi.com:80")
	if err != nil {
		fmt.Println(err)
	}
	req := "GET / HTTP/1.0\r\n" +
		"Host: rickandmortyapi.com\r\n" +
		"User-Agent: Client\r\n\r\n"

	_, err = con.Write([]byte(req))
	if err != nil {
		fmt.Println(err)
	}

	res, err := ioutil.ReadAll(con)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}
