package main

import (
	"net/http"

	"github.com/ambientself/golearn/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3001", nil)
}
