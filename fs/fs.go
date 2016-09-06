// Copyright 2011 ZHU HAIHUA <kimiazhu@gmail.com>. 
// All rights reserved.
// Use of this source code is governed by MIT 
// license that can be found in the LICENSE file.

// Description: fs
// Author: ZHU HAIHUA
// Since: 2016-09-06 00:47
package fs

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kimiazhu/tinyserver/model"
)

func ServeStatic(root string) func(*gin.Context) {
	return func(c *gin.Context) {
		fs := http.FileServer(http.Dir(root))
		if model.Config.Mode == model.DOWNLOAD_MODE {
			c.Writer.Header().Add("Content-Type", "application/octet-stream")
		}
		fs.ServeHTTP(c.Writer, c.Request)
	}
}


