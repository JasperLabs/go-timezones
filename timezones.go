// Package timezones provides an slice of timezones as well as a simple
// API for accessing and searching timezones.
package timezones

import (
	"errors"
)

// TimeZone represents an international date/time zone.
type TimeZone struct {
	Value  string   `json:"value"`
	Abbr   string   `json:"abbr"`
	Offset float32  `json:"offset"`
	Isdst  bool     `json:"isdst"`
	Text   string   `json:"text"`
	Utc    []string `json:"utc"`
}

var (
	TimeZoneNotFound = errors.New("time zone not found")
	NoTimeZonesFound = errors.New("no time zones were found")
)

// GetTimeZones returns a slice of TimeZone.
func GetTimeZones() []TimeZone {
	return goData
}

// GetTimeZoneByValue returns a TimeZone by the Value field.
func GetTimeZoneByValue(value string) (*TimeZone, error) {
	var timeZone *TimeZone
	timeZones := GetTimeZones()

	for i := range timeZones {
		if timeZones[i].Value == value {
			timeZone = &timeZones[i]
			break
		}
	}

	if timeZone == nil {
		return nil, TimeZoneNotFound
	}

	return timeZone, nil
}

// GetTimeZoneByOffset returns all TimeZones where the Offset field equals
// the offset argument.
func GetTimeZoneByOffset(offset float32) ([]*TimeZone, error) {
	var matchingTimeZones []*TimeZone
	timeZones := GetTimeZones()

	for i := range timeZones {
		if timeZones[i].Offset == offset {
			matchingTimeZones = append(matchingTimeZones, &timeZones[i])
		}
	}

	if len(timeZones) == 0 {
		return nil, TimeZoneNotFound
	}

	return matchingTimeZones, nil
}
