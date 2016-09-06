// Author: ZHU HAIHUA
// Date: 9/6/16
package model

type ServerMode int

const (
	DOWNLOAD_MODE ServerMode = iota
	CONTENT_MODE
)

var Config ServerConfig = ServerConfig{Mode: DOWNLOAD_MODE}

type ServerConfig struct {
	Mode ServerMode
}


