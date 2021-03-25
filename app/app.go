package app

import (
	"github.com/joho/godotenv"
	"golang-template/app/routes"
	"golang-template/app/utils"
)

func Run() {
	godotenv.Load()
	utils.InitLogger()
	routes.Run()
}
