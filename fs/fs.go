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
	"github.com/kimiazhu/log4go"
	"strings"
	"net/url"
	"fmt"
	"path/filepath"
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

func ServeFile(root string) func(*gin.Context) {
	return func(c *gin.Context) {
		path := c.Param("filepath")
		path = filepath.Join(".", path)
		up, isRoot := upDir(path)
		log4go.Debug("up dir is: %v, %v", up, isRoot)
		if !isRoot && isDir(path) {
			upUrl := url.URL{Path: up}
			fmt.Fprintf(c.Writer, "<a href=\"/%s\">..</a>\n", upUrl.String())
		}
		http.ServeFile(c.Writer, c.Request, path)
	}
}

func trimDash(path string) string {
	if strings.HasPrefix(path, ".") {
		path = path[1:]
	}

	if strings.HasSuffix(path, "/") {
		path = path[0:len(path) - 1]
	}

	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	return path
}

func upDir(path string) (upPath string, isRoot bool) {
	log4go.Debug("current dir is: %v", path)
	path = trimDash(path)
	lastDash := strings.LastIndex(path, "/")
	if lastDash > 0 {
		path = path[0:lastDash]
		if isDir(path) {
			path = path + "/"
		}
		log4go.Debug("Up dir is: %v", path)
		return path, false
	} else if len(path) > 0 {
		return "", false
	} else {
		log4go.Debug("current dir is root")
		return "", true

	}
}

func isDir(path string) bool {
	if !strings.HasPrefix(path, "/") && !strings.HasPrefix(path, "./") {
		path = "./" + path
	}
	f, err := http.Dir(path).Open("")
	if err != nil {
		log4go.Error("failed to open path: %v", path)
		return false
	}

	d, err := f.Stat()
	if err != nil {
		log4go.Error("failed to get stat of file: %v", path)
		return false
	}

	return d.IsDir()
}
