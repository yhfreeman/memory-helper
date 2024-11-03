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

// long-term memory: 	12hr, 1d, 3d, 7d, 14d, 30d 
func generateEvents(event Event) []Event {
	slots := []int64{
		12 * 60 * 60,
		86400,
		86400*3,
		86400*7,
		86400*14,
		86400*30,
		86400*90,
	}
	var events []Event
	for _, addseconds := range slots {
		tmp := event
		tmp.Tiptime = tmp.Addtime + addseconds
		events = append(events, tmp)
	}
	return events
}
