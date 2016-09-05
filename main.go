// Author: ZHU HAIHUA
// Date: 9/5/16
package main

import (
	"github.com/kimiazhu/ginweb"
	"github.com/kimiazhu/ginweb/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	g := ginweb.New()

	g.Use(func(c *gin.Context){
		c.H
	}).Static("/", "")

	g.StaticFile("/", "");

	ginweb.Run(conf.Conf.SERVER.PORT, g)
}