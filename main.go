package main

import (
	"time"

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
	client, _ := gorm.Open("mysql", "root:root@tcp([mysql]:3306)/test?charset=utf8mb4&parseTime=true")

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

func (client *Client) getRedis(cc *gin.Context) {
	defer client.Redis.Close()

	res_get := Get("sample-key", client.Redis)
	cc.JSON(200, res_get)
}
func (client *Client) getMysql(cc *gin.Context) {
	var data TestData
	client.Mysql.Where("id = 1").Find(&data)
	cc.JSON(200, data)
}

func main() {

	r := gin.Default()

	client := init_client()
	time.Sleep(20)
	client.setup()

	r.GET("/redis", client.getRedis)

	r.GET("/mysql", client.getMysql)
	r.Run()
}
func (client *Client) setup() {

	client.Mysql.AutoMigrate(&TestData{})

	testdata := TestData{Data: "hugahuga", Token: "hunngaaaaaaaa"}
	Set("sample-key", "sample-value", client.Redis)
	client.Mysql.Create(&testdata)

}

type TestData struct {
	gorm.Model
	Data  string `json:"data"`
	Token string `json:"token"`
}
