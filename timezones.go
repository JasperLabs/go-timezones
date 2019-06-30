// Package timezones provides an array of timezones as well as a simple API for accessing
// and searching timezones.
package timezones

import (
	"encoding/json"
	"errors"
)

// TimeZone represents an international date/time zone.
type TimeZone struct {
	Value  string   `json:"value"`
	Abbr   string   `json:"abbr"`
	Offset int      `json:"offset"`
	Isdst  bool     `json:"isdst"`
	Text   string   `json:"text"`
	Utc    []string `json:"utc"`
}

var TimeZoneNotFound = errors.New("timezone not found")

// GetTimeZones returns an array of TimeZone.
func GetTimeZones() []TimeZone {
	var timezones []TimeZone
	json.Unmarshal([]byte(data), &timezones)
	return timezones
}

// GetTimeZoneByValue returns a TimeZone by the Value field.
func GetTimeZoneByValue(value string) (*TimeZone, error) {
	var timezone TimeZone
	timezones := GetTimeZones()

	for i := range timezones {
		if timezones[i].Value == value {
			timezone = timezones[i]
			break
		}
	}

	if len(timezone.Value) == 0 {
		return nil, TimeZoneNotFound
	}

	return &timezone, nil
}
