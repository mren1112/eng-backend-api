package main

import (
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

var ctx = context.Background()

func init() {
	environments.TimeZoneInit()
	environments.EnvironmentInit()
	oracle_db = databases.NewDatabases().OracleInit()
	// redis_cache = databases.NewDatabases().RedisInint()
}

type FOO struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	defer oracle_db.Close()
	// defer redis_cache.Close()
	gin.SetMode(gin.ReleaseMode)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	
	
	routers.Setup(router, oracle_db)
	//router.Run(PORT)
}
