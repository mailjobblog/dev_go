package main

import (
	"fmt"
)

func main() {
	app, _ := InitApp()
	fmt.Println(app.Msg, app.Ber)
}
