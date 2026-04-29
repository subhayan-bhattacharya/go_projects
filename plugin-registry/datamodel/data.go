package datamodel

import "time"

type Sanitizable interface {
	GetId() string
	GetValue() float64
	SetTags(map[string]string)
	GetTags() map[string]string
}

type SensorData struct {
	SensorID  string
	Value     float64
	Unit      string
	TimeStamp time.Time
	Tags      map[string]string
}

func (s *SensorData) GetId() string {
	return s.SensorID
}

func (s *SensorData) GetValue() float64 {
	return s.Value
}

func (s *SensorData) SetTags(tags map[string]string) {
	s.Tags = tags
}

func (s *SensorData) GetTags() map[string]string {
	return s.Tags
}
