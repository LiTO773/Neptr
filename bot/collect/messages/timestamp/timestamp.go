package timestamp

import "regexp"

// GetDayAndHour Returns the hour and day from the timestamp
func GetDayAndHour(timestamp string) (string, string) {
	r := regexp.MustCompile(`([0-9]{4}-[0-9]{2}-[0-9]{2})T([0-9]{2}):`)
	response := r.FindAllStringSubmatch(timestamp, -1)
	return response[0][1], response[0][2]
}
