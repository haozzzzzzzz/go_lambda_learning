package main

import (
	"ExampleXRay/api"

	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
)

func main() {
	engine := ginbuilder.GetEngine()
	api.BindRouters(engine)
	engine.Run(":8100")
}
