package anglebetweenhandsofaclock

import "math"

func angleClock(hour, minutes int) float64 {
	var res float64
	angelMinutes := minutes * 6
	angelHour := float64(hour*30) + float64(minutes)/2.0
	res = math.Abs(angelHour - float64(angelMinutes))
	if res > 180 {
		res = 360 - res
	}

	return res
}
