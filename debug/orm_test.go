package debug

import (
	"assistantor/model"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name   string `gorm:"index" json:"name"`
	UserID string
	// 使用默认外键在写入User的时候默认会把User表中的ID字段写入到Company表中的UserID字段
	// 可使用foreignKey,references指定关联关系, foreignKey:UserName;references:name 会将User表中的name字段写到Company表中的UserName字段
	// foreignKey是其他表字段, references是本表字段
	Cid uint
}
type Company struct {
	gorm.Model
	Name string `json:"name"`
	CSn  string `json:"cSn"`
}

func TestGormFilePartition(t *testing.T) {

	engine, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Info().Msgf("create db error: %v", err)
		return
	}
	engine.AutoMigrate(&model.FilePartition{}, &model.PartitionInfo{})

	p1 := model.PartitionInfo{
		SegmentId:   "segment1",
		SegmentName: "segment1",
		SegmentPath: "zzzzz",
	}

	p2 := model.PartitionInfo{
		SegmentId:   "segment2",
		SegmentName: "segment2",
		SegmentPath: "xxxx",
	}
	f := model.FilePartition{
		FileId:   "1111",
		FileName: "1111.mp4",
		FilePath: "xxxxx/xxxx",
		PartitionList: []model.PartitionInfo{
			p1,
			p2,
		},
	}
	engine.Save(&f)
}

func TestGorm(t *testing.T) {

	engine, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		log.Info().Msgf("create db error: %v", err)
		return
	}
	engine.AutoMigrate(&User{}, &Company{})
	//c := Company{
	//	Name: "cnm",
	//}
	//c.ID = 1
	//c2 := Company{
	//	Name: "cnm2",
	//}
	//c2.ID = 2
	//
	//u := User{
	//	Name: "lance",
	//	Cid: c.ID,
	//}
	//
	//engine.Save(&c)
	//engine.Save(&c2)
	//engine.Save(&u)


	var u User
	err = engine.Table("companies").Select("users.*").Joins("left join users on companies.id = users.cid").Where("companies.id = ?", 1).Find(&u).Error
	log.Info().Msgf("user is: %v, error is: %v", u, err)
}
