package main

import (
	"github.com/gin-gonic/gin"
)

/*
Goal:
- starts an http server at: localhost:3000
- user call http://localhost:3000?name=jessy then server response "received:jessy"
*/

func main() {
	engine := gin.New()
	engine.GET("/", func(g *gin.Context) {
		name := g.Query("name")
		g.String(200, "received:"+name)
	})
	panic(engine.Run(":3000"))
}
