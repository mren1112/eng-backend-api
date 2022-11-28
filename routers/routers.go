package routers

import (
	"fmt"
	"net/http"

	//call repo
	"eng-backend-api/repositories/countercontrolrepo"
	"eng-backend-api/repositories/facultyRepo"

	//call handles
	"eng-backend-api/handlers/countercontrolhandle"
	"eng-backend-api/handlers/facultyHandle"

	//call service
	"eng-backend-api/services/countercontrolservice"
	"eng-backend-api/services/facultyServices"

	//middleware
	"eng-backend-api/middlewares"

	// lib require
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func Setup(router *gin.Engine, oracle_db *sqlx.DB) {

	router.Use(middlewares.NewCorsAccessControll().CorsAccessControll())

	router.GET("/healthz", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status":  "200",
			"message": "The service works normally.",
		})
	})

	facultyRepo := facultyRepo.NewFacultyAllRepo(oracle_db)
	facultyServices := facultyServices.NewFacultyAllService(facultyRepo)
	facultyHandle := facultyHandle.NewFacultyAllHandler(facultyServices)

	countercontrolRepo := countercontrolrepo.NewCounterControlAllRepo(oracle_db)
	countercontrolService := countercontrolservice.NewCounterControlAllService(countercontrolRepo)
	counterControlHandle := countercontrolhandle.NewCounterControlAllHandler(countercontrolService)

	router.GET("/facs", facultyHandle.GetFacultyAll)
	//router.POST("/fac/:param1", facultyHandle.GetFacultyById)
	router.POST("/fac/:param1", facultyHandle.GetFacultyById)
	router.POST("/fx", facultyHandle.GetFacultyByJS)

	router.GET("/counter", counterControlHandle.GetCounterControlAll)

	fmt.Println("hello", viper.GetString("engApi.port"))

	PORT := viper.GetString("engApi.port")
	router.Run(PORT)
}
