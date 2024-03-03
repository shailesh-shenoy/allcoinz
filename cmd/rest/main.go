package main

import (
	"github.com/shailesh-shenoy/allcoinz/api"
)

func main() {

	s := api.ApiServer{
		ListenAddr: ":8080",
	}

	s.Run()
}
