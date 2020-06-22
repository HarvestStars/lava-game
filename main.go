package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lava-game/rpc"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var port string

func init() {
	config()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/lava", rpc.ReadHandler)
	r.GET("/static/images/:addr", rpc.ImageHandler)
	r.GET("/order", rpc.OrderHandler)
	r.GET("/liquid/:slotindex", rpc.LiquidHandler)
	r.GET("/participate", rpc.ParticipateHandler)
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

	rpc.DeadLine = int(viper.Get("backend.deadline").(float64))
	// redis 配置
	rpc.RedisPort = viper.Get("redis.port").(string)
	rpc.RedisType = viper.Get("redis.type").(string)
	rpc.RedisIP = viper.Get("redis.ip").(string)
	rpc.RedisPWD = viper.Get("redis.password").(string)
}
