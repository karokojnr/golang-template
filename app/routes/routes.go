package routes

import (
	"github.com/gin-gonic/gin"
	"golang-template/app/controllers"
	"golang-template/app/utils"
	"golang-template/app/utils/database"
	"log"
	"net/http"
)

var (
	app controllers.App
)

func init() {
	db, err := database.Connect()
	if err != nil {
		utils.LogError("Cannot connect to DB!")
	}
	app.DB = db
	database.AutoMigrateDB(app.DB)
	//Init redis
	database.InitRedis("1")
}


func Run() {
	r := gin.Default()

	// Session to use in authorization
	//r.Use(sessions.Sessions("golang-template-session", sessions.NewCookieStore([]byte(utils.GoDotEnvVariable("SESSION_KEY")))))
	// Serve static files
	//r.Static("/assets", utils.GoDotEnvVariable("ASSETS_FOLDER"))
	// Load html templates
	//r.LoadHTMLGlob(utils.GoDotEnvVariable("TEMPLATES_FOLDER"))

	r.GET("/", app.GetIndex)
	r.GET("/cars", app.CreateCar)
	r.POST("/cars", app.CreateCar)
	r.GET("/cars/:id", app.FindCar)
	r.PATCH("/cars/:id", app.UpdateCar)
	r.DELETE("/cars/:id", app.DeleteBook)
	r.NoRoute(app.NotFound)

	port := utils.GetPort()
	utils.Log("App running on port  ✓  :", port)
	log.Fatal(http.ListenAndServe(port, r))
}
