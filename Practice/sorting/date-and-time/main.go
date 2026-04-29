package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type DurationWrapper time.Duration

func (d *DurationWrapper) UnmarshalJSON(b []byte) error {
	unquotedString := strings.Trim(string(b), `"`)
	parsedValue, err := time.ParseDuration(unquotedString)
	if err != nil {
		return fmt.Errorf("Could not parse the string into duration %w", err)
	}
	*d = DurationWrapper(parsedValue)
	return nil
}

func (d DurationWrapper) String() string {
	return time.Duration(d).String()
}

type Event struct {
	Name string `json:"name"`
	StartUpTime time.Time `json:"start_time"`
	Duration DurationWrapper `json:"duration"`
	Priority int `json:"priority"`
}

func (e Event) Endtime() time.Time {
	return e.StartUpTime.Add(time.Duration(e.Duration))
}

func DoesEventsOverlap(event1 , event2 Event) bool {
	if event1.StartUpTime.Before(event2.Endtime()) && event2.StartUpTime.Before(event1.Endtime()) {
		return true
	}
	return false
}

func compareEventByStartUpTime(event1, event2 Event) int {
	return event1.StartUpTime.Compare(event2.StartUpTime)
}

func compareEventByOverlapAndPriority(event1, event2 Event) int {
	if DoesEventsOverlap(event1, event2) {
		return cmp.Compare(event2.Priority, event1.Priority)
	}
	return compareEventByStartUpTime(event1, event2)
}

type Events struct {
	Events []Event `json:"events"`
}


// func compareEventByEndTime(event1, event2 Event) int {
// 	return event1.Endtime().Compare(event2.Endtime())
// }

func getData(inputFile string) (Events, error) {
	var events Events
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return events, fmt.Errorf("Could not read file %w", err)
	}
	err = json.Unmarshal(data, &events)
	if err != nil {
		return events, fmt.Errorf("Could not unmarshall contents %w", err)
	}
	return events, nil
}

func main() {
	jsonFileData, err := getData("data.json")
	if err != nil {
		fmt.Printf("We have fucked up %v\n", err)
		os.Exit(1)
	}
	// slices.SortFunc(jsonFileData.Events, compareEventByStartUpTime)
	slices.SortFunc(jsonFileData.Events, compareEventByOverlapAndPriority)
	for _, event := range jsonFileData.Events {
		fmt.Println(event)
	}
}
