package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testcases = []struct {
	name        string
	inputDate1  time.Time
	inputDate2  time.Time
	expected    int
	holidayList map[time.Time]bool
}{
	{"sameDayWorkDay", time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC), 1, map[time.Time]bool{}},
	{"sameDayHoliDay", time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC), 0, map[time.Time]bool{}},
	{"startDateHoliday", time.Date(2023, 8, 19, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 21, 0, 0, 0, 0, time.UTC), 1, map[time.Time]bool{}},
	{"endDateHoliday", time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 9, 3, 0, 0, 0, 0, time.UTC), 0, map[time.Time]bool{}},
	{"multipleHolidaysInRow", time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC), 0, map[time.Time]bool{
		time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 21, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 22, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 23, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 24, 0, 0, 0, 0, time.UTC): true,
		time.Date(2023, 8, 18, 0, 0, 0, 0, time.UTC): true,
	}},
	{"noRemaining", time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC), 10, map[time.Time]bool{}},
	{"holidaysInBetween", time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Date(2023, 8, 11, 0, 0, 0, 0, time.UTC), 7, map[time.Time]bool{
		time.Date(2023, 8, 9, 0, 0, 0, 0, time.UTC):   true,
		time.Date(2023, 9, 16, 0, 0, 0, 0, time.UTC):  true, //odd man
		time.Date(2023, 8, 14, 0, 0, 0, 0, time.UTC):  true, //odd man
		time.Date(2023, 10, 14, 0, 0, 0, 0, time.UTC): true, //odd man
		time.Date(2023, 8, 14, 0, 0, 0, 0, time.UTC):  true, //odd man
	}},
}

func TestCountWorkDay(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, countWorkday(tc.inputDate1, tc.inputDate2, tc.holidayList))
			// // arrange
			// expected := tc.expected
			// // act
			// got := countWorkday(tc.inputDate1, tc.inputDate2, tc.holidayList)
			// // assert
			// if expected != got {
			// 	t.Errorf("expected : %d got : %d ", expected, got)
			// }
		})
	}
}

// func Test_countWorkday(t *testing.T) {
// 	type args struct {
// 		inputDate1  time.Time
// 		inputDate2  time.Time
// 		holidayList map[time.Time]bool
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := countWorkday(tt.args.inputDate1, tt.args.inputDate2, tt.args.holidayList); got != tt.want {
// 				t.Errorf("countWorkday() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
