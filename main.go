package main

import (
	"receipt-processor/routes"
)

func main() {
	r := routes.Router()

	r.SetTrustedProxies(nil)
	r.Run()
}
