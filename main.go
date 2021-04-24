package main

import (
	"net/http"

	"github.com/mahajany/gowebservice/models/controllers"
)

const (
	PORT = ":3000"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(PORT, nil)
}
