package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"api2go-gin-gorm-simple/model"
	"api2go-gin-gorm-simple/resource"
	"api2go-gin-gorm-simple/storage"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gingonic"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver("/"),
		gingonic.New(r),
	)

	db, err := storage.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userStorage := storage.NewUserStorage(db)
	api.AddResource(model.User{}, resource.UserResource{UserStorage: userStorage})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Run(":31418")
}
