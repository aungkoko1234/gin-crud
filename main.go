package main

import (
	"net/http"
	"time"

	"github.com/aungkoko1234/gin-crud/config"
	"github.com/aungkoko1234/gin-crud/controller"
	"github.com/aungkoko1234/gin-crud/model"
	"github.com/aungkoko1234/gin-crud/repository"
	"github.com/aungkoko1234/gin-crud/router"
	"github.com/aungkoko1234/gin-crud/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	//db
    db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//init repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)

	//init service
	tagsService := service.NewTagServiceImpl(tagsRepository, validate)

	//init controller 
	tagsController := controller.NewTagController(tagsService)

	//init router

	routes := router.NewRouter(tagsController)



	server := &http.Server{
		Addr:  ":8888",
		Handler: routes,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}