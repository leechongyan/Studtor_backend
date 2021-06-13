package timeslot_helpers

import "strconv"

const (
	originTime int = 8
)

func convertHoursToString(hour int) (time string) {
	if hour < 12 {
		time = strconv.Itoa(hour) + "am"
	} else {
		if hour > 12 {
			hour -= 12
		}
		time = strconv.Itoa(hour) + "pm"
	}
	return
}

func ConvertSlotIdToTimeString(slotId int) (time string) {
	startTime := originTime + slotId
	startTimeString := convertHoursToString(startTime)
	endTimeString := convertHoursToString(startTime + 1)
	return startTimeString + " to " + endTimeString
}
