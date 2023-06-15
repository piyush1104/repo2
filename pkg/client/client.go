package client

import (
	"fmt"

	"github.com/piyush1104/repo2/pkg"
)

func GetApp() {
	// get data and return app
	fmt.Println(pkg.App{ID: "hello"})
}
