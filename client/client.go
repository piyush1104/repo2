package client

import (
	interfaces "example.com/interfaces"
)

func GetApp() {
	// get data and return app
	return interfaces.App{ID: "hello"}
}
