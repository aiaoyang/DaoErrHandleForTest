package main

import (
	"log"

	"github.com/aiaoyang/errhandle/pkg/service"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	service.Query("")
}
