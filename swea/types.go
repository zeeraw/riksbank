package swea

import (
	"strconv"
	"time"
)

// Language is the container for a language ID
type Language string

const (
	// English is the english language ID
	English = Language("en")

	// Swedish is the swedish language ID
	Swedish = Language("sv")
)

// SearchGroupSeries represents searchable group series
type SearchGroupSeries struct {
	GroupID  string
	SeriesID string
}

// CrossPair are the series to compare in a currency exchange
type CrossPair struct {
	BaseSeriesID    string
	CounterSeriesID string
}

// CrossRateInfo is an exchange rate between two currencies
type CrossRateInfo struct {
	Base    string
	Counter string
	Date    time.Time
	Period  string
	Value   string
	Average string
}

// DayInfo represents a date in the context of the central bank
type DayInfo struct {
	Date      time.Time
	Week      int
	WeekYear  int
	IsBankDay bool
}

// CrossSeriesInfo represents a interest or currency conversion series information
type CrossSeriesInfo struct {
	ID          string
	Name        string
	Description string
}

// SeriesInfo represents a interest or currency conversion series information
type SeriesInfo struct {
	ID              string
	GroupID         string
	Name            string
	Description     string
	LongDescription string
	Source          string
	Type            string
	From            time.Time
	To              time.Time
}

// RateInfo represents information about a rate for a series in a period
type RateInfo struct {
	GroupID   string
	GroupName string

	SeriesID   string
	SeriesName string

	Date   time.Time
	Period string

	Average string
	Min     string
	Max     string
	Ultimo  string
	Value   string
}

// GroupInfo represents a grouping of interest or exchange rates
type GroupInfo struct {
	ID       string
	ParentID string

	Name        string
	Description string
}

// GroupsInfo represents several groups group info
type GroupsInfo []GroupInfo

func (gis GroupsInfo) Len() int {
	return len(gis)
}

func (gis GroupsInfo) Swap(i, j int) {
	gis[i], gis[j] = gis[j], gis[i]
}

func (gis GroupsInfo) Less(i, j int) bool {
	a, _ := strconv.Atoi(gis[i].ID)
	b, _ := strconv.Atoi(gis[j].ID)
	return a < b
}

// GetCalendarDaysRequest represents the parameters to get all business days between two dates
type GetCalendarDaysRequest struct {
	From time.Time
	To   time.Time
}

// GetCalendarDaysResponse contains the
type GetCalendarDaysResponse struct {
	From time.Time
	To   time.Time
	Days []DayInfo
}

// GetAllCrossNamesRequest represents the parameters get all the exchange rate series suitable for cross rate names
type GetAllCrossNamesRequest struct {
	Language Language
}

// GetAllCrossNamesResponse contains the currency conversion series
type GetAllCrossNamesResponse struct {
	Language Language
	Series   []CrossSeriesInfo
}

// GetCrossRatesRequest represents the parameters to get all change rates
type GetCrossRatesRequest struct {
	CrossPairs []CrossPair

	From            time.Time
	To              time.Time
	Language        Language
	AggregateMethod string
}

// GetCrossRatesResponse contains exchange rates
type GetCrossRatesResponse struct {
	CrossRates []CrossRateInfo
	CrossPairs []CrossPair

	From            time.Time
	To              time.Time
	Language        Language
	AggregateMethod string
}

// GetInterestAndExchangeRatesRequest represents the parameters to get exchange and interest rates
type GetInterestAndExchangeRatesRequest struct {
	Series []SearchGroupSeries

	From            time.Time
	To              time.Time
	Language        Language
	AggregateMethod string

	Average bool
	Min     bool
	Max     bool
	Ultimo  bool
}

// GetInterestAndExchangeRatesResponse contains interest and exchange rates
type GetInterestAndExchangeRatesResponse struct {
	Rates []RateInfo

	Series []SearchGroupSeries

	From            time.Time
	To              time.Time
	Language        Language
	AggregateMethod string

	Average bool
	Min     bool
	Max     bool
	Ultimo  bool
}

// GetInterestAndExchangeGroupNamesRequest represents the parameters to get a list of all groups
type GetInterestAndExchangeGroupNamesRequest struct {
	Language Language
}

// GetInterestAndExchangeGroupNamesResponse contains all groups
type GetInterestAndExchangeGroupNamesResponse struct {
	Groups   []GroupInfo
	Language Language
}

// GetInterestAndExchangeNamesRequest represents the parameters to get all series for a group
type GetInterestAndExchangeNamesRequest struct {
	GroupID  string
	Language Language
}

// GetInterestAndExchangeNamesResponse contains all series for a group
type GetInterestAndExchangeNamesResponse struct {
	Series   []SeriesInfo
	GroupID  string
	Language Language
}
