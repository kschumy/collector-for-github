package types

import (
	"fmt"
	"time"
)

type RelativeTimeProvider interface {
	GetRelativeTime() RelativeTime
}

type RelativeTime struct {
	// What relation (before or after) dateTime has to what is being queried.
	relative Relative
	// Time in UTC
	dateTime time.Time
}

func InitializeWithDefault() *RelativeTime {
	return &RelativeTime{
		relative: AfterDateTime,
		dateTime: getEarliestDate(),
	}
}

func NewRelativeTime(relative Relative, dateTime time.Time) (*RelativeTime, error) {
	convertedTime := ConvertToUTC(&dateTime)
	err := checkRelativeAndDateTimeCombo(relative, convertedTime)
	if err != nil {
		return nil, err
	}

	return &RelativeTime{
		relative: relative,
		dateTime: ConvertToUTC(&convertedTime),
	}, nil
}

func (rt RelativeTime) GetRelative() Relative {
	return rt.relative
}

func (rt RelativeTime) GetTime() time.Time {
	return rt.dateTime
}

func (rt *RelativeTime) SetTime(dateTime time.Time) error {
	convertedTime := ConvertToUTC(&dateTime)
	err := checkRelativeAndDateTimeCombo(rt.relative, convertedTime)
	if err != nil {
		return err
	}
	rt.dateTime = convertedTime
	return nil
}

func (rt *RelativeTime) SetRelative(relative Relative) error {
	err := checkRelativeAndDateTimeCombo(relative, rt.dateTime)
	if err != nil {
		return err
	}
	rt.relative = relative
	return nil
}

func GetCopyOrDefault(rt RelativeTime) (*RelativeTime, error) {
	if rt.GetRelative() == AnyDateTime || rt.GetTime().IsZero() {
		return InitializeWithDefault(), nil
	}
	return NewRelativeTime(rt.GetRelative(), rt.GetTime())
}

func (rt *RelativeTime) ErrorIfInvalid() error {
	return checkRelativeAndDateTimeCombo(rt.relative, rt.dateTime)
}

func (rt *RelativeTime) IsValid() bool {
	return checkRelativeAndDateTimeCombo(rt.relative, rt.dateTime) == nil
}

// BUG: fix invalid combos
func checkRelativeAndDateTimeCombo(relative Relative, dateTime time.Time) error {
	if !relative.IsValid() {
		return fmt.Errorf("invalid relative: %v", relative)
	}

	//if hasCorrectLocation(dateTime) {
	//	return fmt.Errorf("invalid time location: %v", dateTime.Zone())
	//}

	if relative == AnyDateTime && !dateTime.IsZero() {
		return fmt.Errorf("cannot provide a dateTime if relative is AnyTime")
	}

	if relative != AnyDateTime && dateTime.IsZero() {
		return fmt.Errorf("must provide a dateTime if relative is not AnyTime")
	}
	// TODO: invalid if in the future and AfterDateTime

	return nil
}

func ConvertToUTC(t *time.Time) time.Time {
	newTime := *t
	if !hasCorrectLocation(t) {
		_, secondsToOffset := t.Zone()
		newTime = t.Add(time.Second * time.Duration(secondsToOffset))
	}
	return newTime.UTC()
}

// Returns 'true' if t is in UTC and 'false' otherwise.
func hasCorrectLocation(t *time.Time) bool {
	return t.Location() == time.UTC
}

func getEarliestDate() time.Time {
	return time.Date(2008, 2, 5, 9, 0, 0, 0, time.UTC)
}
