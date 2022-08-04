package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"tag-engine/pkg/global"
)

func Init(configPath string) *gin.Engine {
	initConfig(configPath)
	initDB()
	engine := initEngine()
	initRouter(engine)
	return engine
}

func initConfig(configPath string) {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("初始化配置失败", err)
	}
}

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.name"))
	gormConfig := &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("初始化数据库失败", err)
	}
	global.MysqlClient = db
}

func initEngine() *gin.Engine {
	engine := gin.Default()
	gin.SetMode(viper.GetString("server.run_mod"))
	return engine
}

func Run(engine *gin.Engine) {
	err := engine.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.Fatal("启动服务失败", err)
	}
}
