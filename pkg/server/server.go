package server

import (
	"fmt"

	util "github.com/piyush1104/repo3"

	"github.com/piyush1104/repo2/pkg"
)

func main() {
	x := pkg.App{ID: util.GetID()}
	fmt.Println(x)
}

func GetApp() {
	x := pkg.App{ID: util.GetID()}
	fmt.Println(x)
}
