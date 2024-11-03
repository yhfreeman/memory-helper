package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strconv"
	"fmt"
)

func addEvent(event Event) (bool, error) {
    storage, err := LoadEvents(serverConfigs.dbFile)
    if err != nil {
        return false, err
    }

    // Add the new event
    storage.Events = append(storage.Events, event)

    // Save to YAML file
    if err := SaveEvents(serverConfigs.dbFile, storage); err != nil {
        fmt.Printf("Error saving events: %v\n", err) // Log error
        return false, err
    }

    return true, nil
}

func getEvents(owner string, starttime string, endtime string) ([]Event, error) {
	storage, err := LoadEvents(serverConfigs.dbFile)
	if err != nil {
		fmt.Printf("Error loading events: %v\n", err) // Log the loading error
		return nil, err
	}

	fmt.Printf("PATHs: %v\n", serverConfigs.dbFile) // Log file path if loading is successful
	var ret []Event 
	for _, event := range storage.Events {
		if event.Owner == owner && (event.Addtime >= parseTime(starttime)) && (event.Addtime <= parseTime(endtime)) {
			ret = append(ret, event)
		}
	}
	return ret, nil
}


// Helper function to parse time string
func parseTime(timeStr string) int64 {
	// Parse the time string to int64, handle error as needed
	// Assume the input is valid for simplicity
	value, _ := strconv.ParseInt(timeStr, 10, 64)
	return value
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
		Owner := "yfreeman"
		events, err := getEvents(Owner, starttime, endtime)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
		} else {
			context.JSON(http.StatusOK, gin.H{"ret": events})
		}
	})
	
	app.GET("/addevent", func(context *gin.Context) {
		title := context.DefaultQuery("title", "no title")
		description := context.DefaultQuery("description", "no description")
		Owner := "yfreeman"
		event := Event{
			Type:        "test",
			Title:       title,
			Description: description,
			Addtime:     time.Now().Unix(),
			Tiptime:     time.Now().Unix() + 10000,
			Owner:       Owner,
		}
		
		// Log the event being added
		fmt.Printf("Adding event: %+v\n", event) // Log event details
		
		isSuccess, err := addEvent(event)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
		} else {
			context.JSON(http.StatusOK, gin.H{"ret": isSuccess})
		}
	})
	

	app.Run(serverConfigs.addr)
}
