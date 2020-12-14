package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Client struct {
	Redis redis.Conn
	Mysql *gorm.DB
}

func init_client() Client {
	client, _ := gorm.Open("mysql", "root:root@tcp([mysql]:3306)/Weather?charset=utf8mb4&parseTime=true")

	const Addr = "redis:6379"

	c, _ := redis.Dial("tcp", Addr)

	return Client{Mysql: client, Redis: c}

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

func (client *Client) hoge(cc *gin.Context) {
	defer client.Redis.Close()

	res_set := Set("sample-key", "sample-value", client.Redis)
	fmt.Println(res_set) // OK

	res_get := Get("sample-key", client.Redis)
	fmt.Println(res_get) // sample-value

	cc.JSON(200, res_get)

}
func main() {

	r := gin.Default()

	client := init_client()

	r.GET("/", client.hoge)
	r.Run() // デフォルトが8080ポートなので今回は変更しません
}
