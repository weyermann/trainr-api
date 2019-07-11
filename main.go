package main
 
import (
	// "github.com/gorilla/handlers"
	"github.com/weyermann/trainr-api/app"
	"github.com/weyermann/trainr-api/config"
)
 
func main() {
	config := config.GetConfig()

	// handlers.AllowedOrigins([]string{"*"})
 
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
