package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Connection() redis.Conn {
	const Addr = "redis:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

func Set(key, value string, c redis.Conn) string {
	res, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		panic(err)
	}
	return res
}

func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}

func hoge(cc *gin.Context) {
	c := Connection()
	defer c.Close()

	res_set := Set("sample-key", "sample-value", c)
	fmt.Println(res_set) // OK

	res_get := Get("sample-key", c)
	fmt.Println(res_get) // sample-value

	cc.JSON(200, res_get)

}
func main() {
	r := gin.Default()
	// setup()
	r.GET("/", hoge)
	r.Run() // デフォルトが8080ポートなので今回は変更しません
}
