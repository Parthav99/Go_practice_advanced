package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// parameterized Testing
var testcases = []struct {
	name         string
	businessDays int
	inputTime    time.Time
	expected     time.Time
	holidayList  map[time.Time]bool
}{
	{"oneValue", 1, time.Date(2023, 9, 9, 0, 0, 0, 0, time.UTC), time.Date(2023, 9, 11, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"sundayStart", 1, time.Date(2023, 8, 6, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 7, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"zeroValue", 0, time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), map[time.Time]bool{}},
	{"negativeValue", -25, time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC), time.Date(2023, 7, 26, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 13, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 14, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC): true,
	}},
	{"holidaysInBetween", 11, time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 17, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC):  true,
		time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 16, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
	}},
	{"holidays", 4, time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 9, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
		time.Date(2023, 8, 3, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 8, 0, 0, 0, 0, time.UTC): true,
	}},
}

func TestCalculateEndDate(t *testing.T) {
	for _, testCase := range testcases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, calculateEndDate(testCase.inputTime, testCase.businessDays, testCase.holidayList))
		})
	}
}

//----------------------------------------------------------------------//
// {"multipleHolidays", 5,
//  time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC),
//  time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC), map[time.Time]bool{
// 	time.Date(2023, 8, 25, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 8, 28, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 8, 30, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 8, 31, 0, 0, 0, 0, time.UTC): true,
// 	time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC):  true,
// 	time.Date(2023, 9, 4, 0, 0, 0, 0, time.UTC):  true,
// 	time.Date(2023, 9, 5, 0, 0, 0, 0, time.UTC):  true,
// 	time.Date(2023, 9, 6, 0, 0, 0, 0, time.UTC):  true,
// 	time.Date(2023, 9, 7, 0, 0, 0, 0, time.UTC):  true,
// 	time.Date(2023, 9, 8, 0, 0, 0, 0, time.UTC):  true,
// }},

// //Arrange
// expected := testCase.expected
// //Act
// got := CalculateEndDate(testCase.inputTime, testCase.businessDays, testCase.holidayList)
// if got != expected {
// 	t.Errorf("expected : %s got : %s", expected.Format("2006-01-02"), got.Format("2006-01-02"))
// }
