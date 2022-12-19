package main

import (
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tamahavik/tasker-api/pkg/config"
)

func main() {
	err := ReadConfiguration()
	if err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("MODE"))
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".pdf", ".mp4"})))

	conn := config.NewConnection()
	db, err := conn.GetConnection()
	if err != nil {
		panic(err)
	}

	config.NewRoutes(r, db).InitRoute()

	err1 := r.Run(fmt.Sprintf(":%s", viper.GetString("PORT")))
	if err != nil {
		panic(err1)
	}
}
