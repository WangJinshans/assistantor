package model

import "gorm.io/gorm"

type FilePartition struct {
	gorm.Model
	FileId        string          `gorm:"index" json:"file_id,omitempty"`
	FileName      string          `json:"file_name,omitempty"`
	FilePath      string          `json:"file_path,omitempty"`
	PartitionList []PartitionInfo `gorm:"foreignKey:FileId;references:file_id" json:"partition_list,omitempty"`
}

type PartitionInfo struct {
	gorm.Model
	FileId      string `gorm:"size:50"`
	SegmentId   string `json:"segment_id"`
	SegmentName string `json:"segment_name"`
	SegmentPath string `json:"segment_path"`
}
