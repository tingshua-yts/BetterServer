package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("../frontend/my-app/build/", true)))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 11
		} else {
			count = v.(int)
			count++
		}
		fmt.Println("set value:")
		fmt.Println(count)

		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})

	})
	r.Run(":8000")
}
