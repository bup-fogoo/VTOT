package main

import (
	"VTOT/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 2)
	router := gin.Default()
	// 加载 静态资源
	router.Static("/tmp", "./tmp/")
	router.LoadHTMLGlob("view/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.MaxMultipartMemory = 8 << 20 // 8 MiB, default 32 MiB

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		//文件后缀校验
		extString := strings.ToLower(path.Ext(file.Filename))
		//允许上传文件的格式
		allowExtMap := map[string]bool{
			".mp4": true,
			".mp3": true,
			".wav": true,
		}
		if _, ok := allowExtMap[extString]; !ok {
			c.String(http.StatusBadRequest, "upload file type err")
			return
		}

		basePath := "./tmp/"
		filename := basePath + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("file %s upload to /tmp success ", file.Filename))

		// 处理视频，video to mp3
		vtotmp3 := ""
		if extString == ".mp4" {
			vtotmp3 += service.VideoToAudioService(filename)
		}
		/*   worker  */
		// 音频结果输出格式
		tmpFileName := uuid.New().ID()
		outFilePath := fmt.Sprintf("./tmp/%d.html", tmpFileName)
		c.String(http.StatusMovedPermanently, fmt.Sprintf("\r\n\r\n转文字已成功请访问下面跟目录（需要等几分钟）\r\n/tmp/%d.html\r\n\r\n\r\n如果没有请刷新网页或者重新上传", tmpFileName))

		fileLink := "https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee583045a.wav"
		//fileLink := fmt.Sprintf("http://114.116.37.179/upload/%s", vtotmp3)
		wg.Add(1)
		go service.Worker(fileLink, outFilePath, ch, &wg)

	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
