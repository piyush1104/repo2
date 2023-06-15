package server

import (
	"fmt"

	util "github.com/piyush1104/repo3"

	"github.com/piyush1104/repo2"
)

func main() {
	x := repo2.App{ID: util.GetID()}
	fmt.Println(x)
}

func GetApp() {
	x := repo2.App{ID: util.GetID()}
	fmt.Println(x)
}
