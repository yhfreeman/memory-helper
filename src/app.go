package main

import (
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"strconv"
	"fmt"
)

func createStudyItem(studyItemToAdd StudyItem) (bool, error) {
    // Load existing data from the YAML file
    storage, err := LoadDB(serverConfigs.dbFile)
    if err != nil {
        fmt.Printf("Error loading storage: %v\n", err)
        return false, err
    }

    // Append the new StudyItem to storage
    storage.StudyItems = append(storage.StudyItems, studyItemToAdd)

    // Generate events based on the studyItem creation
    generatedEvents := generateEvents(studyItemToAdd)

    // Append each generated event to storage.Events
    storage.Events = append(storage.Events, generatedEvents...)

    // Save to YAML file
    if err := WriteDB(serverConfigs.dbFile, storage); err != nil {
        fmt.Printf("Error saving events: %v\n", err) // Log error
        return false, err
    }

    return true, nil
}

func getEvents(owner string, starttime string, endtime string) ([]Event, error) {
	storage, err := LoadDB(serverConfigs.dbFile)
	if err != nil {
		fmt.Printf("Error loading events: %v\n", err) // Log the loading error
		return nil, err
	}

	fmt.Printf("PATHs: %v\n", serverConfigs.dbFile) // Log file path if loading is successful
	var ret []Event 

	// Create a map to quickly access StudyItems by UID
	studyItemMap := make(map[string]StudyItem)
	for _, studyItem := range storage.StudyItems {
		studyItemMap[studyItem.ID] = studyItem
	}

	// Iterate through events and check conditions
	for _, event := range storage.Events {
		// Find the corresponding StudyItem for the current Event
		studyItem, exists := studyItemMap[event.StudyItemID] // Assuming StudyItemID in Event references the StudyItem UID
		if !exists {
			continue // If StudyItem does not exist, skip this event
		}

		// Check if the event's associated StudyItem matches the owner and falls within the time range
		if studyItem.Owner == owner && (studyItem.Addtime >= parseTime(starttime)) && (studyItem.Addtime <= parseTime(endtime)) {
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

	app.GET("/getevents", func(context *gin.Context) {
		starttime := context.DefaultQuery("starttime", "0")
		endtime := context.DefaultQuery("endtime", "9999999999")
		owner := "yfreeman"
		events, err := getEvents(owner, starttime, endtime)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
		} else {
			context.JSON(http.StatusOK, gin.H{"ret": events})
		}
	})
	
	app.GET("/addstudyitem", func(context *gin.Context) {
		title := context.DefaultQuery("title", "no title")
		description := context.DefaultQuery("description", "no description")
		owner := "yfreeman"

		studyItem := StudyItem{
			ID:          uuid.New().String(),
			Title:       title,
			Description: description,
			Addtime:     time.Now().Unix(),
			Owner:       owner,
		}

		fmt.Printf("Adding study item: %+v\n", studyItem)
		
		isStudyItemAdded, err := createStudyItem(studyItem)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"ret": err})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"ret": isStudyItemAdded})
		}
	})
	

	app.Run(serverConfigs.addr)
}
