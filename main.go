package main

import (
	"VTOT/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	InitConfig()
	router := gin.Default()
	// 加载 静态资源
	router.Static("/tmp", "./tmp/")
	router.LoadHTMLGlob("view/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.MaxMultipartMemory = 8 << 20 // 8 MiB, default 32 MiB

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("fileUpload")
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
		filename := filepath.Join(basePath, filepath.Base(file.Filename))
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		// 处理视频，video to mp3
		if extString == ".mp4" {
			file.Filename = service.VideoToAudioService(filename)
		}
		/*
			worker
			可以存自己的服务器，可以存阿里云空间，我这里是存的自己的服务器
		*/
		//fileLink := "https://gw.alipayobjects.com/os/bmw-prod/0574ee2e-f494-45a5-820f-63aee583045a.wav"
		fileLink := fmt.Sprintf("%stmp/%s", viper.GetString("vps"), file.Filename)
		res := service.Worker(fileLink)
		if res == nil {
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		c.JSON(http.StatusOK, res)

	})
	// 启动HTTP服务,默认在0.0.0.0:8080启动服务
	port := viper.GetString("server.port")
	if port != "" {
		panic(router.Run(":" + port))
	}
	panic(router.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
