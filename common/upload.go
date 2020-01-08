package common

import (
	"awesomeProject/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 上传文件
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	path := "./static/upload/" + time.Now().Format("20060102")
	filename := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename

	// 文件夹不存在就创建
	if !exists(path) {
		if err := os.MkdirAll(path, os.ModePerm);err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				Msg:   "文件夹创建失败",
			})
			return
		}
	}

	// 保存上传的文件
	if err := c.SaveUploadedFile(file, path + "/" +  filename); err != nil{
		c.JSON(http.StatusOK, serializer.Response{Msg:"文件上传失败",})
		return
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code: 1,
		Data: path[1 : len(path)] + "/" + filename,
		Msg:  "上传成功",
	})
}


// 判断是否存在文件夹
func exists(path string) bool  {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}