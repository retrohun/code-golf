package hole

import "strconv"

var timeUnits = [...]struct {
	seconds          int
	singular, plural string
}{
	{60 * 60 * 24 * 365 * 1000, "a millenium", "millenia"},
	{60 * 60 * 24 * 365, "a year", "years"},
	{60 * 60 * 24 * 30, "a month", "months"},
	{60 * 60 * 24 * 7, "a week", "weeks"},
	{60 * 60 * 24, "a day", "days"},
	{60 * 60, "an hour", "hours"},
	{60, "a minute", "minutes"},
	{1, "a second", "seconds"},
}

func formatDistance(secs int) string {
	if secs == 0 {
		return "now"
	}
	past := secs < 0
	if past {
		secs = -secs
	}
	result := ""
	for _, v := range timeUnits {
		if v.seconds <= secs {
			q := secs / v.seconds
			if q == 1 {
				result = v.singular
			} else {
				result = strconv.Itoa(q) + " " + v.plural
			}
			break
		}
	}
	if past {
		return result + " ago"
	}
	return "in " + result
}

var _ = answerFunc("time-distance", func() []Answer {
	inputs := []int{0}

	timeUnitsChosen := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i <= 32; i++ {
		// Randomize which timeUnits will appear.
		timeUnitsChosen = append(timeUnitsChosen, randInt(1, 7))
	}

	for _, j := range timeUnitsChosen {
		secs := timeUnits[j].seconds
		secsLarger := timeUnits[j-1].seconds
		inputs = append(inputs, randInt(secs, secs*2-1))        // future singular
		inputs = append(inputs, -randInt(secs, secs*2-1))       // past singular
		inputs = append(inputs, randInt(2*secs, secsLarger-1))  // future plural
		inputs = append(inputs, -randInt(2*secs, secsLarger-1)) // past plural
		inputs = append(inputs, 2*secs)                         // future exactly 2
		inputs = append(inputs, -2*secs)                        // past exactly 2
		blimit := min(secs-1, 1000)
		a := randInt(2, 6)
		b := randInt(-blimit, blimit)
		inputs = append(inputs, a*secs+b) // future plural antiapproximation
		a = -randInt(2, 6)
		b = randInt(-blimit, blimit)
		inputs = append(inputs, a*secs+b) // past plural antiapproximation
	}

	tests := make([]test, len(inputs))
	for i, inp := range inputs {
		tests[i] = test{strconv.Itoa(inp), formatDistance(inp)}
	}

	return outputTests(shuffle(tests))
})
