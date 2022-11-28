package main

import (
<<<<<<< HEAD
	"context" 

	//call db

	"eng-backend-api/databases"
	"eng-backend-api/environments"


	//call roting
	"eng-backend-api/routers"

	// lib require
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx" 
)

var oracle_db *sqlx.DB
=======
	"context"
	"eng-backend-api/databases"
	"eng-backend-api/environments"
	"eng-backend-api/handlers/employeehandle"
	"eng-backend-api/handlers/universityhandle"
	"eng-backend-api/services/employeeservices"
	"eng-backend-api/services/universityservices"

	"eng-backend-api/handlers/studentshandle"
	"eng-backend-api/repositories/employeerepo"
	"eng-backend-api/repositories/studentsrepo"
	"eng-backend-api/repositories/universityrepo"
	"eng-backend-api/services/studentservices"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var oracle_db *sqlx.DB
var redis_cache *redis.Client
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d

var ctx = context.Background()

func init() {
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	// redis_cache = databases.NewDatabases().RedisInint()
}

<<<<<<< HEAD
type FOO struct {
	URL string `json:"url" binding:"required"`
}

=======
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
func main() {
	defer oracle_db.Close()
	// defer redis_cache.Close()
	gin.SetMode(gin.ReleaseMode)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

<<<<<<< HEAD
	
	
	routers.Setup(router, oracle_db)
	//router.Run(PORT)
=======
	employeeRepo := employeerepo.NewEmployeeAllRepo(oracle_db)
	employeeService := employeeservices.NewEmployeeAllService(employeeRepo)
	employeeHandle := employeehandle.NewEmployeeAllHandler(employeeService)

	router.GET("/employee", employeeHandle.GetEmployeeAllUnicon)

	universityRepo := universityrepo.NewUniversityAllRepo(oracle_db)
	universityService := universityservices.NewUniversityAllService(universityRepo)
	universityHandle := universityhandle.NewUniversityAllHandler(universityService)

	router.GET("/university", universityHandle.GetUniversityAllUnicon)

	studentRoute := router.Group("/student")
	{

		studentRepo := studentsrepo.NewStudentAllRepo(oracle_db)
		studentService := studentservices.NewStudentAllService(studentRepo)
		studentHandle := studentshandle.NewStudentAllHandler(studentService)
		studentRoute.GET("/:param1", studentHandle.GetStudentLimitUnicon)
		studentRoute.GET("/", studentHandle.GetStudentAllUnicon)

	}

	PORT := viper.GetString("ruUnicon.port")
	router.Run(PORT)
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
}
