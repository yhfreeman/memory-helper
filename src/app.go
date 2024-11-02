package main

import (
	"github.com/gin-gonic/gin"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
	"net/http"
)

const KEY_PREFIX = "zset:"

func addEvent(event Event, redisConfigs RedisConfigs) (bool, error){
	client, err := redis.Dial(redisConfigs.network, redisConfigs.address)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	eventstr := Encode(event)
	score := event.Addtime
	if score == 0 {
		event.Addtime = time.Now().Unix()
	}
	key := KEY_PREFIX + event.Owner
	events := generateEvents(event)
	for _, event := range events {
		eventstr = Encode(event)
		_, err = client.Do("zadd", key, event.Tiptime, eventstr)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func getEvents(redisConfigs RedisConfigs, Owner string, starttime string, endtime string)  ([]Event, error) {
	client, err := redis.Dial(redisConfigs.network, redisConfigs.address)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	key := KEY_PREFIX + Owner
	content, err := redis.Values(client.Do("zrangebyscore", key, starttime, endtime))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var ret []Event
	for _, v := range content {
		//fmt.Printf("%s ", v.([]byte))
		tmp, _ := Decode(string(v.([]byte)))
		ret = append(ret, tmp)
	}
	return ret, nil
}


func main() {
	app := gin.Default()

	app.Static("/static", serverConfigs.static)


	app.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "is healthy")
	})

	app.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	app.GET("/add_entry", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_entry.html", nil)
	})

	app.GET("/getevent", func(context *gin.Context) {
		starttime := context.DefaultQuery("starttime", "0")
		endtime := context.DefaultQuery("endtime", "9999999999")
		fmt.Println(starttime, endtime)
		Owner := "yfreeman"
		events, err := getEvents(redisConfigs, Owner, starttime, endtime)
		fmt.Println(events)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err, "starttime":starttime, "endtime":endtime})
		}else{
			context.JSON(http.StatusOK, gin.H{"ret": events, "starttime":starttime, "endtime":endtime})
		}
	})

	app.GET("/addevent", func(context *gin.Context) {
		title := context.DefaultQuery("title", "no title")
		description := context.DefaultQuery("description", "no description")
		Owner := "yfreeman"
		event := Event{
			Type:"test",
			Title:title,
			Description:description,
			Addtime:time.Now().Unix(),
			Tiptime:time.Now().Unix() + 10000,
			Owner:Owner,
		}
		isSuccess, err := addEvent(event, redisConfigs)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
		}else{
			context.JSON(http.StatusOK, gin.H{"ret": isSuccess})
		}
	})

	app.Run(serverConfigs.addr)
}
