package main

import (
	"gadget-points/infrastructure/auth"
	"gadget-points/infrastructure/persistence"
	"gadget-points/interfaces"
	"gadget-points/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// redis details
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	services, err := persistence.NewRepositories(dbDriver, user, password, port, host, dbname)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func(services *persistence.Repositories) {
		err := services.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(services)

	err = services.Automigrate()
	if err != nil {
		return
	}

	redisService, err := auth.NewRedisDB(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()
	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	agent := interfaces.NewAgent(services.Agent, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)
	order := interfaces.NewOrder(services.Order, services.Product, services.Agent, services.Activity, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// user routes
	r.POST("/users", users.SaveUser)
	r.GET("/users", users.GetUsers)
	r.GET("/users/:user_id", users.GetUser)

	// authentication routes
	r.POST("/login", authenticate.Login)
	r.POST("/logout", authenticate.Logout)
	r.POST("/refresh", authenticate.Refresh)

	// agent routes
	r.GET("/agent/:agent_code", agent.GetAgent)

	// order routes
	r.POST("/order", order.CreateOrder)

	// Other routes
	// TODO

	// Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8888"
	}

	log.Fatal(r.Run(":" + appPort))
}
