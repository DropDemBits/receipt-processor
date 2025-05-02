package main

import (
	"receipt-processor/routes"
)

func main() {
	r := routes.Router()

	r.Run()
}
