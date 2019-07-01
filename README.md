# Timezones for Go

`go get github.com/JasperLabs/go-timezones`

Available on GoDoc at [godoc.org/github.com/JasperLabs/go-timezones](https://godoc.org/github.com/JasperLabs/go-timezones).

Provides a Go application with a slice of timezones (sourced from [dmfilipenko/timezones.json](https://github.com/dmfilipenko/timezones.json/blob/master/timezones.json)) as package `timezones`.

## Usage

Access TimeZone struct via `timezones.TimeZone`
```go
type TimeZone struct {
	Value  string   `json:"value"`  // Dateline Standard Time
	Abbr   string   `json:"abbr"`   // DST
	Offset int      `json:"offset"` // -12
	Isdst  bool     `json:"isdst"`  // false
	Text   string   `json:"text"`   // (UTC-12:00) International Date Line West
	Utc    []string `json:"utc"`    // ["Etc/GMT+12"]
}
```

Get a slice of TimeZones.
```go
timezones := GetTimeZones()
	
for _, tz := range timezones {
	log.Printf("%v", tz.Text)
}
```

Get a specific TimeZone by `Value`.
```go
var offset string
tz, err := timezones.GetTimeZoneByValue("New Zealand Standard Time")
if err == timezones.TimeZoneNotFound {
	return nil, err
}
offset = tz.Offset
```

Get a slice of TimeZones by `Offset`.
```go
tzs, err := timezones.GetTimeZoneByOffset(12)
if err == timezones.NoTimeZonesFound {
	return nil, err
}

for _, tz := range tzs {
	log.Printf("%v", tz.Text)
}
```

## Contributing
Pull requests with tests are welcome and will be reviewed in a timely manner (24-hour window).

## License

The MIT License (MIT)

Copyright (c) 2019 [Jasper Property Limited](https://www.jasper.io).

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
