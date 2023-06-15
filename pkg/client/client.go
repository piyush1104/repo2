package client

import (
	"fmt"

	"github.com/piyush1104/repo2"
)

func GetApp() {
	// get data and return app
	fmt.Println(repo2.App{ID: "hello"})
}
