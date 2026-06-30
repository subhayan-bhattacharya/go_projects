package observer

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

type Payload struct {
	Timestamp   time.Time
	StationID   string
	Temperature float64 // in Celsius
	Humidity    float64 // percentage
	AQI         int     // Air Quality Index
}

type Observer interface {
	PushPayload(payload Payload) error
}

type Publisher interface {
	Notify() error
	Subscribe(observer Observer) error
	UnSubscribe(observer Observer) error
}

type WeatherStation struct {
	lock        sync.Mutex
	data        Payload
	subscribers []Observer
}

func (w *WeatherStation) Notify() error {
	//This part is important where we shrink the critical section so that
	//we do not hold the entire lock for the entire duration of push payload
	//It is possible that push payload again calls unsubscribe , which would deadlock
	w.lock.Lock()
	subscribers := slices.Clone(w.subscribers)
	w.lock.Unlock()
	for _, observer := range subscribers {
		err := observer.PushPayload(w.data)
		if err != nil {
			fmt.Printf("could not push payload to struct %+v, error is %v", observer, err)
		}
	}
	return nil
}

func (w *WeatherStation) Subscribe(observer Observer) error {
	if observer == nil {
		return errors.New("cannot subscribe a nil observer")
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	w.subscribers = append(w.subscribers, observer)
	return nil
}

func (w *WeatherStation) UnSubscribe(observer Observer) error {
	if observer == nil {
		return errors.New("cannot unsubscribe a nil observer")
	}
	w.lock.Lock()
	defer w.lock.Unlock()
	idx := slices.IndexFunc(w.subscribers, func(target Observer) bool {
		return target == observer
	})
	if idx != -1 {
		w.subscribers = slices.Delete(w.subscribers, idx, idx+1)
	} else {
		return errors.New("there is no such subscriber to unsubscribe")
	}
	return nil
}

type Receiver struct{}

func (r *Receiver) PushPayload(data Payload) error {
	fmt.Printf("received payload %+v", data)
	return nil
}
