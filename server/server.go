package server

import (
	"fmt"

	"github.com/piyush1104/interfaces"
	util "github.com/piyush1104/repo3"
)

func main() {
	x := interfaces.App{ID: util.GetID()}
	fmt.Println(x)
}

func GetApp() {
	x := interfaces.App{ID: util.GetID()}
	fmt.Println(x)
}
