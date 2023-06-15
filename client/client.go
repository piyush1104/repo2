package client

import (
	"fmt"

	"github.com/piyush1104/interfaces"
)

func GetApp() {
	// get data and return app
	fmt.Println(interfaces.App{ID: "hello"})
}
