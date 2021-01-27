package main

import (
	"github.com/ramailh/auth-server/http/router"
	"github.com/ramailh/auth-server/props"
	"log"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	if err := gotenv.Load(env); err != nil {
		log.Println(err)
	}

	props.Setup()
}
func main() {
	rtr := router.NewHTTPRouter()
	log.Fatal(rtr.Run(":" + props.Port))
}
