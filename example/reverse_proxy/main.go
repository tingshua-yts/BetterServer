package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {

		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

func main() {
	r := gin.Default()
	proxy, _ := NewProxy("http://39.107.37.181:3000")

	group := r.Group("/api/v1/pipeline")
	group.GET("/*path", func(c *gin.Context) {
		// path := c.Param("path")
		// fmt.Println("path:" + path)

		fmt.Println("before url:" + c.Request.URL.Path)
		c.Request.URL.Path = strings.Replace(c.Request.URL.Path, "/api/v1/pipeline", "", 1)
		fmt.Println("after url:" + c.Request.URL.Path)
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		proxy.ServeHTTP(c.Writer, c.Request)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
