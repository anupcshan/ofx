package main

import (
	"log"
	"os"

	"github.com/anupcshan/ofx"
)

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	fx, err := ofx.Parse(os.Stdin)

	if err != nil {
		panic(err)
	}

	log.Println(fx)
}
