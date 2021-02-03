package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	Week       = "WEEK"
	Month      = "MONTH"
	Quater     = "QUARTER"
	Year       = "YEAR"
	LastSunday = "LAST_SUNDAY_OF_YEAR"

	DateFormat = "2006-01-02"
	Space      = " "
)

var (
	Location = time.Now().UTC().Location()
)

type Interval struct {
	StartDate time.Time
	EndDate   time.Time
}

func main() {

	var (
		intervalType string
		str          string
	)
	_, err := fmt.Fscan(os.Stdin, &intervalType)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscan(os.Stdin, &str)
	if err != nil {
		log.Fatal(err)
	}

	begin, err := time.Parse(DateFormat, str)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscan(os.Stdin, &str)
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse(DateFormat, str)
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	intertvals, err := SplitDateInterval(begin, end, intervalType)
	if err != nil {
		log.Fatal(err)
	}
	stop := time.Now()
	fmt.Println(len(intertvals))

	for i := range intertvals {
		fmt.Println(
			intertvals[i].StartDate.Format(DateFormat),
			intertvals[i].EndDate.Format(DateFormat),
		)
	}

	fmt.Printf("Elapsed time: %v", stop.Sub(start))
}

func SplitDateInterval(begin, end time.Time, interval string) ([]Interval, error) {
	if begin.After(end) {
		return nil, fmt.Errorf("begin > end")
	}

	switch interval {
	case Week:
		return splitDateIntervalByWeek(begin, end)
	case Month:
		return splitDateIntervalByMonth(begin, end)
	case Quater:
		return splitDateIntervalByQuater(begin, end)
	case Year:
		return splitDateIntervalByYear(begin, end)
	case LastSunday:
		return splitDateIntervalByLastSunday(begin, end)
	default:
		return nil, fmt.Errorf("No such interval type: %s", interval)
	}
}

func splitDateIntervalByWeek(begin, end time.Time) ([]Interval, error) {
	intervals := make([]Interval, 0)

	var firstSunday time.Time

	switch begin.Weekday() {
	case time.Monday:
		firstSunday = begin.AddDate(0, 0, 6)
	case time.Tuesday:
		firstSunday = begin.AddDate(0, 0, 5)
	case time.Wednesday:
		firstSunday = begin.AddDate(0, 0, 4)
	case time.Thursday:
		firstSunday = begin.AddDate(0, 0, 3)
	case time.Friday:
		firstSunday = begin.AddDate(0, 0, 2)
	case time.Saturday:
		firstSunday = begin.AddDate(0, 0, 1)
	case time.Sunday:
		firstSunday = begin
	}

	if !firstSunday.Before(end) {
		intervals = append(intervals, Interval{
			StartDate: begin,
			EndDate:   end,
		})
		return intervals, nil
	}

	intervals = append(intervals, Interval{
		StartDate: begin,
		EndDate:   firstSunday,
	})

	for sunday := firstSunday; sunday.Before(end); {
		monday := sunday.AddDate(0, 0, 1) // Monday
		if monday.After(end) {
			break
		}
		if monday.Equal(end) {
			intervals = append(intervals, Interval{
				StartDate: monday,
				EndDate:   monday,
			})
			break
		}
		nextSunday := sunday.AddDate(0, 0, 7) // Sunday
		if !nextSunday.Before(end) {
			intervals = append(intervals, Interval{
				StartDate: monday,
				EndDate:   end,
			})
			break
		}
		intervals = append(intervals, Interval{
			StartDate: monday,
			EndDate:   nextSunday,
		})

		sunday = nextSunday // Sunday = next Sunday
	}
	return intervals, nil
}

func splitDateIntervalByYear(begin, end time.Time) ([]Interval, error) {
	intervals := make([]Interval, 0)

	beginYear := begin.Year()
	endYear := end.Year()

	for year := beginYear; year <= endYear; year++ {
		i := Interval{}

		if year == beginYear {
			i.StartDate = begin
		} else {
			i.StartDate = time.Date(year, 01, 01, 0, 0, 0, 0, Location)
		}
		lastDayOfYear := time.Date(year, 12, 31, 0, 0, 0, 0, Location)
		if lastDayOfYear.After(end) {
			i.EndDate = end
		} else {
			i.EndDate = lastDayOfYear
		}
		intervals = append(intervals, i)
	}

	return intervals, nil
}

func splitDateIntervalByQuater(begin, end time.Time) ([]Interval, error) {
	intervals := make([]Interval, 0)

	years, err := splitDateIntervalByYear(begin, end)
	if err != nil {
		return nil, err
	}

	for y := range years {
		interval, err := splitYearByQuater(years[y].StartDate, years[y].EndDate)
		if err != nil {
			return nil, err
		}
		intervals = append(intervals, interval...)
	}

	return intervals, nil
}

func splitYearByQuater(begin, end time.Time) ([]Interval, error) {
	if begin.Year() != end.Year() {
		return nil, fmt.Errorf("not same year")
	}
	if !begin.Before(end) {
		return nil, fmt.Errorf("end < begin")
	}
	intervals := make([]Interval, 0)

	switch begin.Month() {
	case time.January, time.February, time.March:
		switch end.Month() {
		case time.January, time.February, time.March:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   end,
				})
			return intervals, nil
		case time.April, time.May, time.June:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 03, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 04, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		case time.July, time.August, time.September:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 03, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 04, 01, 0, 0, 0, 0, Location),
					EndDate:   time.Date(begin.Year(), 06, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 07, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		case time.October, time.November, time.December:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 03, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 04, 01, 0, 0, 0, 0, Location),
					EndDate:   time.Date(begin.Year(), 06, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 07, 01, 0, 0, 0, 0, Location),
					EndDate:   time.Date(begin.Year(), 9, 30, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 10, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		}
	case time.April, time.May, time.June:
		switch end.Month() {
		case time.April, time.May, time.June:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   end,
				})
			return intervals, nil
		case time.July, time.August, time.September:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 06, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 07, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		case time.October, time.November, time.December:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 06, 31, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 07, 01, 0, 0, 0, 0, Location),
					EndDate:   time.Date(begin.Year(), 9, 30, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 10, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		}
	case time.July, time.August, time.September:
		switch end.Month() {
		case time.July, time.August, time.September:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   end,
				})
			return intervals, nil
		case time.October, time.November, time.December:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   time.Date(begin.Year(), 9, 30, 0, 0, 0, 0, Location),
				},
				Interval{
					StartDate: time.Date(begin.Year(), 10, 01, 0, 0, 0, 0, Location),
					EndDate:   end,
				},
			)
			return intervals, nil
		}
	case time.October, time.November, time.December:
		switch end.Month() {
		case time.October, time.November, time.December:
			intervals = append(intervals,
				Interval{
					StartDate: begin,
					EndDate:   end,
				})
			return intervals, nil
		}
	}
	return intervals, nil
}

func splitDateIntervalByMonth(begin, end time.Time) ([]Interval, error) {
	intervals := make([]Interval, 0)

	years, err := splitDateIntervalByYear(begin, end)
	if err != nil {
		return nil, err
	}

	for y := range years {
		interval, err := splitYearByMonth(years[y].StartDate, years[y].EndDate)
		if err != nil {
			return nil, err
		}
		intervals = append(intervals, interval...)
	}

	return intervals, nil
}

func IsYear366(y int) bool {
	if y%4 != 0 {
		return false
	} else if y%100 == 0 {
		if y%400 == 0 {
			return true
		}
		return false
	} else {
		return true
	}
}

func splitYearByMonth(begin, end time.Time) ([]Interval, error) {
	if begin.Year() != end.Year() {
		return nil, fmt.Errorf("not same year")
	}
	// if !begin.Before(end) {
	// 	return nil, fmt.Errorf("end < begin")
	// }
	// intervals := make([]Interval, 0)
	// days := 28
	// if IsYear366(begin.Year()) {
	// 	days++
	// }

	// JanuaryStart := time.Date(begin.Year(), 01, 1, 0, 0, 0, 0, Location)
	// JanuaryEnd := time.Date(begin.Year(), 01, 31, 0, 0, 0, 0, Location)

	// FebStart := time.Date(begin.Year(), 02, 1, 0, 0, 0, 0, Location)
	// FebEnd := time.Date(begin.Year(), 02, days, 0, 0, 0, 0, Location)

	// MarchStart := time.Date(begin.Year(), 03, 1, 0, 0, 0, 0, Location)
	// MarchEnd := time.Date(begin.Year(), 03, 31, 0, 0, 0, 0, Location)

	// AprilStart := time.Date(begin.Year(), 04, 1, 0, 0, 0, 0, Location)
	// AprilEnd := time.Date(begin.Year(), 04, 30, 0, 0, 0, 0, Location)

	// MayStart := time.Date(begin.Year(), 05, 1, 0, 0, 0, 0, Location)
	// MayEnd := time.Date(begin.Year(), 05, 31, 0, 0, 0, 0, Location)

	// JuneStart := time.Date(begin.Year(), 06, 1, 0, 0, 0, 0, Location)
	// JuneEnd := time.Date(begin.Year(), 06, 30, 0, 0, 0, 0, Location)

	// JuleStart := time.Date(begin.Year(), 07, 1, 0, 0, 0, 0, Location)
	// JuleEnd := time.Date(begin.Year(), 07, 31, 0, 0, 0, 0, Location)

	// AugustStart := time.Date(begin.Year(), 8, 1, 0, 0, 0, 0, Location)
	// AugustEnd := time.Date(begin.Year(), 8, 31, 0, 0, 0, 0, Location)

	// SeptemberStart := time.Date(begin.Year(), 9, 1, 0, 0, 0, 0, Location)
	// SeptemberEnd := time.Date(begin.Year(), 9, 30, 0, 0, 0, 0, Location)

	// OctoberStart := time.Date(begin.Year(), 10, 1, 0, 0, 0, 0, Location)
	// OctoberEnd := time.Date(begin.Year(), 10, 31, 0, 0, 0, 0, Location)

	// NovemberStart := time.Date(begin.Year(), 11, 1, 0, 0, 0, 0, Location)
	// NovemberEnd := time.Date(begin.Year(), 11, 30, 0, 0, 0, 0, Location)

	// DecemberStart := time.Date(begin.Year(), 12, 1, 0, 0, 0, 0, Location)
	// DecemberEnd := time.Date(begin.Year(), 12, 31, 0, 0, 0, 0, Location)
	return nil, nil
}

func splitDateIntervalByLastSunday(begin, end time.Time) ([]Interval, error) {
	return nil, nil
}
