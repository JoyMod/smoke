package main

import (
	"net/http"
	"smoke/smoke"
)

func main() {

	s := smoke.New()

	v1 := s.Group("/v1")
	{

		v1.GET("/hello", func(c *smoke.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Query("name"), c.Path)
		})

		v1.GET("/hello/:name", func(c *smoke.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
		})

		v1.GET("/assets/*filepath", func(c *smoke.Context) {
			c.JSON(http.StatusOK, smoke.H{
				"filepath": c.Param("filepath"),
			})
		})
	}
	v2 := s.Group("/v2")
	{
		v2.GET("hello/:name", func(c *smoke.Context) {
			c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *smoke.Context) {
			c.JSON(http.StatusOK, smoke.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	s.Run(":8080")
}
