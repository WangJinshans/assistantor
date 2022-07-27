package model

import "gorm.io/gorm"

type FilePartition struct {
	gorm.Model
	FileId        string          `json:"file_id,omitempty"`
	FileName      string          `json:"file_name,omitempty"`
	FilePath      string          `json:"file_path,omitempty"`
	PartitionList []PartitionInfo `json:"partition_list,omitempty"`
}

type PartitionInfo struct {
	gorm.Model
	Id          int64
	FileId      string
	SegmentId   string
	PartitionId int
	FilePath    string
}
