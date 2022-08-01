package debug

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string `gorm:"index" json:"name"`
	// 使用默认外键在写入User的时候默认会把User表中的ID字段写入到Company表中的UserID字段
	// 可使用foreignKey,references指定关联关系, foreignKey:UserName;references:name 会将User表中的name字段写到Company表中的UserName字段
	// foreignKey是其他表字段, references是本表字段
	Company []Company `gorm:"foreignKey:UserName;references:name"`
}
type Company struct {
	gorm.Model
	ID       int    `json:"ID"`
	UserName string `gorm:"size:10"`
	Name     string `json:"name"`
	CSn      string `json:"cSn"`
}

func TestGorm(t *testing.T) {

	engine, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Info().Msgf("create db error: %v", err)
		return
	}
	engine.AutoMigrate(&User{}, &Company{})
	c := Company{
		ID:   1,
		Name: "cnm",
	}
	c2 := Company{
		ID:   2,
		Name: "cnm2",
	}

	u := User{
		Name: "lance",
		Company: []Company{
			c, c2,
		},
	}

	//engine.Save(&c)
	//engine.Save(&c2)
	//engine.Save(&u)

	var user User
	engine.Model(&User{}).Preload("Company").Where("name = ?", "lance").First(&user)
	log.Info().Msgf("user is: %v", user)

	var companyList []Company
	engine.Model(&u).Association("Company").Find(&companyList)
	log.Info().Msgf("company list is: %v", companyList)
}
