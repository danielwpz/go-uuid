package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	uuid := uuid.NewV4()
	fmt.Println(uuid)
}
