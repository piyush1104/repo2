package server

import (
	interfaces "example.com/interfaces"
	util "example.com/repo3"
	"fmt"
)

func main() {
	x := interfaces.App{ID: util.GetID()}
	fmt.Println(x)
}
