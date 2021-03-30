package database

import (
	"fmt"
	redis "github.com/go-redis/redis/v7"
	"golang-template/app/models"
	"golang-template/app/utils"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"strconv"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=karokojnr dbname=cars port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database!")
	}
	return db, nil
}
func AutoMigrateDB(db *gorm.DB) {
	utils.Log("auto-migrations running...")
	db.AutoMigrate(&models.Car{})
	utils.Log("auto-migration complete...")
}

//RedisClient ...
var RedisClient *redis.Client

//InitRedis ...
func InitRedis(params ...string) {

	var redisHost = utils.GoDotEnvVariable("REDIS_HOST")
	var redisPassword = utils.GoDotEnvVariable("REDIS_PASSWORD")

	db, _ := strconv.Atoi(params[0])

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       db,
	})
}

//GetRedis ...
func GetRedis() *redis.Client {
	return RedisClient
}
