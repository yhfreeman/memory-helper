package main

import (
	"fmt"
	"encoding/json"
)

func Decode(content string)(Event, error) {
	bucket := &Event{}
	json.Unmarshal([]byte(content), bucket)
	//fmt.Println(bucket)
	return *bucket, nil
}

func Encode(event Event) string {
	bytes, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func generateEvents(studyItem StudyItem) []Event {
	slots := []int64{
        12 * 60 * 60,    // 12 hours
        86400,           // 1 day
        86400 * 3,       // 3 days
        86400 * 7,       // 7 days
        86400 * 14,      // 14 days
        86400 * 30,      // 30 days
        86400 * 90,      // 90 days
	}
	var events []Event
    for _, offset := range slots {
        event := Event{
            StudyItemID: studyItem.ID, // Link to StudyItem by ID
            Tiptime:     studyItem.Addtime + offset,
            IsCompleted: false,
        }
        events = append(events, event)
    }
    return events
}
