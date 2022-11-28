package databases

import (
	"fmt"
<<<<<<< HEAD
=======

	"github.com/go-redis/redis/v8"
>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
	_ "github.com/godror/godror"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type connection struct{}

func NewDatabases() *connection {
	return &connection{}
}

func (c *connection) OracleInit() *sqlx.DB {
	db, err := oracleConnection()
	if err != nil {
		panic(err)
	}
	return db
}

<<<<<<< HEAD
 
=======
func (c *connection) RedisInint() *redis.Client {
	return redisConnection()
}

func redisConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis_cache.addressInLocal"),
		// Addr:     viper.GetString("redis_cache.addressInContainer"),
		// Addr:     viper.GetString("redis_cache.addressInServer"),
		Password: viper.GetString("redis_cache.password"),
		DB:       viper.GetInt("redis_cache.db-num"),
	})
}

>>>>>>> d59ef4c7f3f3949a5d269274d97355325461518d
func oracleConnection() (*sqlx.DB, error) {

	dns := fmt.Sprintf("%v", viper.GetString("db.connection"))
	driver := viper.GetString("db.openDriver")

	return sqlx.Open(driver, dns)

}
