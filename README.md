TinyServer的目的是提供一个便携的HTTP文件服务器。

目标是在全平台上，将TinyServer放置于任意目录下运行，无需任何配置即可将此目录中的文件通过HTTP协议共享给其他人访问下载。

另外，通过配置文件提供有限但必要的定制功能，包括：

1. 运行模式：
	1. Nomal：将文本文件的一部分浏览器可识别的文件直接在浏览器上展示出来。
	2. Download：浏览器点击所有文件都启用下载。
2. 列举目录：
当访问目录的时候，列举目录中的所有文件。