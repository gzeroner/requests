package main

import (
	"fmt"
	"github.com/gzeroner/requests"
	"log"
)

func main() {
	res, err := requests.Get("https://www.bing.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.String())
}
