package main

import (
    "code.google.com/p/go-uuid/uuid"
	"fmt"
)

func main () {
	s := uuid.New()
	fmt.Println(s)
}