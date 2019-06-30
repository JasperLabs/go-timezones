package timezones

import (
	"testing"
)

// TestGetTimeZones ensures that an array of TimeZone is returned by GetTimeZones
func TestGetTimeZones(t *testing.T) {
	timezones := GetTimeZones()

	if len(timezones) == 0 {
		t.Error("no timezones returned")
	}
}

// TestGetTimeZoneByValue ensures that GetTimeZoneByValue finds a correct TimeZone by its field Value.
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
