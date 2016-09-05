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

	//application/octet-stream
	g.StaticFS("/", gin.Dir("", true))

	ginweb.Run(conf.Conf.SERVER.PORT, g)
}

