package main
 
import (
	"github.com/weyermann/trainr-api/app"
	"github.com/weyermann/trainr-api/config"
)
 
func main() {
	config := config.GetConfig()
 
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
