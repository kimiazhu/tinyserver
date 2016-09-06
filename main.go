// Author: ZHU HAIHUA
// Date: 9/5/16
package main

import (
	"github.com/kimiazhu/ginweb"
	"github.com/kimiazhu/ginweb/conf"
	//"github.com/gin-gonic/gin"
	"github.com/kimiazhu/tinyserver/fs"
)

func main() {
	g := ginweb.New()

	g.GET("/*filepath", fs.ServeStatic(""))
	//g.GET("/", fs.ServeStatic(""))

	ginweb.Run(conf.Conf.SERVER.PORT, g)
}

