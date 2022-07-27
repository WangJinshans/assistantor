package v1

import (
	"assistantor/model"
	"assistantor/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var FileApi ApiFile

type ApiFile struct {
}

func (*ApiFile) UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "file error",
		})
		return
	}
	basePath := "./upload/"
	filename := basePath + filepath.Base(file.Filename)
	if err = ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

// 断点续传
// 初始化分片列表
func (*ApiFile) InitMultiplePart(ctx *gin.Context) {
	var info model.FilePartition
	err := ctx.BindJSON(&info)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "file error",
		})
		return
	}

	repository.InitFilePartitions(info)

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}


func (*ApiFile) UploadPartition(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "file error",
		})
		return
	}
	basePath := "./upload/partition/"
	filename := basePath + filepath.Base(file.Filename)
	if err = ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func (*ApiFile) FinishMultiplePart(ctx *gin.Context) {
	var info model.FilePartition
	err := ctx.BindJSON(&info)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "file error",
		})
		return
	}

	repository.FinishFilePartitions(info)

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}