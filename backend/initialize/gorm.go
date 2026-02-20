package initialize

import (
	"log"

	"github.com/Smash249/xenova/backend/internal/global"
	"github.com/Smash249/xenova/backend/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initGorm() {
	dsn := viper.GetString("Database.mysqlConfig.dsn")
	stringSize := viper.GetUint("Database.mysqlConfig.defaultStringSize")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: stringSize,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	if err := db.AutoMigrate(models.User{}, models.ProductSeries{}, models.Product{}, models.JournalismSeries{}, models.Journalism{}, models.JournalismSeries{}, models.Journalism{}, models.SoftwareSeries{}, models.Software{}, models.CompanyHonor{}, models.CompanyPatent{}, models.LoveActivity{}); err != nil {
		log.Fatal("数据库模型迁移失败", err)
	}
	global.DB = db
}
