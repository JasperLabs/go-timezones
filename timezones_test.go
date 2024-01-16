package timezones

import (
	"testing"
)

// TestGetTimeZones ensures that a slice of TimeZone is returned by GetTimeZones
func TestGetTimeZones(t *testing.T) {
	timezones := GetTimeZones()

	if len(timezones) == 0 {
		t.Error("no timezones returned")
	}
}

// TestGetTimeZoneByValue ensures that GetTimeZoneByValue finds a correct
// TimeZone by its field Value.
func TestGetTimeZoneByValue(t *testing.T) {
	testCases := map[string]string{
		"SA Pacific Standard Time":  "(UTC-05:00) Bogota, Lima, Quito",
		"New Zealand Standard Time": "(UTC+12:00) Auckland, Wellington",
		"Tonga Standard Time":       "(UTC+13:00) Nuku'alofa",
	}

	for key, testCase := range testCases {
		tz, err := GetTimeZoneByValue(key)
		if err != nil {
			t.Error(err)
		}

		if tz.Text != testCase {
			t.Errorf("timezone text unexpected, expected \"%v\" but got \"%v\"", testCase, tz.Text)
		}
	}
}

// TestGetTimeZoneByOffset ensures that TestGetTimeZoneByOffset finds
// all timezones with the same offset as a sample.
func TestGetTimeZoneByOffset(t *testing.T) {
	testCases := map[float32]int{
		// offset: expected count
		-5: 3,
		-1: 2,
		0:  4,
		1:  5,
		5:  4,
	}

	for offset, expectedCount := range testCases {
		timeZones, err := GetTimeZoneByOffset(offset)
		if err != nil {
			t.Error(err)
		}

		if len(timeZones) != expectedCount {
			t.Errorf(
				"invalid number of TimeZones returned for offset %v, expected %v but got %v",
				offset,
				expectedCount,
				len(timeZones),
			)
		}
	}
}
