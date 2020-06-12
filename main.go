package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/lava-game/rpc"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var port string
var redisPort string
var redisType string
var redisIP string
var redisPWD string

func init() {
	config()
}

func main() {
	rpc.RediCon, _ = redis.Dial(redisType, redisIP+":"+redisPort)
	if _, err := rpc.RediCon.Do("AUTH", redisPWD); err != nil {
		rpc.RediCon.Close()
		fmt.Print("redis权限错误!")
	}
	defer rpc.RediCon.Close()

	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/lava", rpc.ReadHandler)
	r.GET("/static/images/:addr", rpc.ImageHandler)

	r.Run(":" + port)
}

func config() {
	viper.SetConfigName("config") // the name of configure file
	viper.AddConfigPath(".")      // path
	viper.SetConfigType("json")   // type of file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}

	rpcPort := viper.Get("rpc.port").(string)
	port = os.Getenv("PORT")
	if port == "" {
		port = rpcPort
		log.Printf("Defaulting to port %s", port)
	}

	// redis 配置
	redisPort = viper.Get("redis.port").(string)
	redisType = viper.Get("redis.type").(string)
	redisIP = viper.Get("redis.ip").(string)
	redisPWD = viper.Get("redis.password").(string)
}
