package repository

import (
	"assistantor/model"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"os"
)

func InitFilePartitions(ps model.FilePartition) {
	engine.Save(&ps)
}

func GetFilePartitions(ps model.FilePartition) (partitionFiles []model.PartitionInfo) {
	engine.Where("file_id  ?", ps.FileId).Find(&partitionFiles)
	return
}

// 合并文件
func FinishFilePartitions(ps model.FilePartition) (err error) {

	filePartitions := GetFilePartitions(ps)
	fileBasePath := "./upload/"
	fileName := fileBasePath + ps.FileId

	var f *os.File
	f, err = os.Create(fileName)
	if err != nil {
		log.Error().Msgf("failed to create file, error is: %v", err)
		return
	}
	defer f.Close()

	for _, item := range filePartitions {
		basePath := "./upload/partition/"
		filename := basePath + item.SegmentPath
		var bs []byte
		bs, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Error().Msgf("read partition file failed, error is: %v", err)
			return
		}
		_, err = f.Write(bs)
		if err != nil {
			log.Error().Msgf("write partition file failed, error is: %v", err)
			return
		}
	}
	ps.FilePath = fileName
	engine.Save(&ps)
	return
}

// 中间层, 内部可以实现oss 七牛 本地等方式
func SaveFile() {

}
