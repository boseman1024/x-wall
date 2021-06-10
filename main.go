package main

import (
	"github.com/gin-contrib/cors"
	"x-wall/router"
)

func main(){
	r := router.NewRouter()
	r.Use(cors.Default())
	r.Run(":3000")
}
