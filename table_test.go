package table

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type Book struct {
	Name   string
	Author string
	URL    string
}

type User struct {
	ID      string
	Name    string
	Age     int
	Deleted bool
	Created time.Time
}

type Number struct {
	Exported   float32
	unexported float32
}

func TestOutput(t *testing.T) {
	Convey("Output content without panic", t, func() {
		s := []Book{
			{"Go by Example", "Mark McGranaghan", "https://github.com/mmcgrana/gobyexample"},
			{"The Little Go Book", "Karl Seguin", "https://github.com/karlseguin/the-little-go-book"},
		}
		Println()
		Output(s)
	})
}

func TestTable(t *testing.T) {
	Convey("Get table content", t, func() {
		Convey("non-slice, should panic", func() {
			So(func() {
				Table(nil)
			}, ShouldPanicWith, "sliceconv: param \"slice\" should be on slice value")
		})

		Convey("slice of int, should panic", func() {
			s := []int{1, 2, 3}
			So(func() {
				Table(s)
			}, ShouldPanicWith, "Table: items of slice should be on struct value")
		})

		Convey("unexported field should be ignored", func() {
			s := []Number{{1.2, 2.4}, {4.8, 9.6}}
			content := Table(s)
			expected := `
┌──────────┐
│ Exported │
├──────────┤
│ 1.2      │
│ 4.8      │
└──────────┘`
			So("\n"+content, ShouldEqual, expected)
		})

		Convey("Field: string, int, boo, time.Time", func() {
			s := []User{
				{"8", "Captain Jack Sparrow", 31, false, time.Date(2017, time.November, 8, 23, 12, 43, 249437302, time.UTC)},
				{"9", "William Turner", 18, false, time.Date(2009, time.February, 9, 2, 4, 25, 363779979, time.UTC)},
				{"10", "Davy Jones", 120, true, time.Date(1965, time.October, 10, 31, 10, 26, 106273532, time.UTC)},
			}
			content := Table(s)
			expected := `
┌────┬──────────────────────┬─────┬─────────┬─────────────────────────────────────────┐
│ ID │ Name                 │ Age │ Deleted │ Created                                 │
├────┼──────────────────────┼─────┼─────────┼─────────────────────────────────────────┤
│ 8  │ Captain Jack Sparrow │ 31  │ false   │ 2017-11-08 23:12:43.249437302 +0000 UTC │
│ 9  │ William Turner       │ 18  │ false   │ 2009-02-09 02:04:25.363779979 +0000 UTC │
│ 10 │ Davy Jones           │ 120 │ true    │ 1965-10-11 07:10:26.106273532 +0000 UTC │
└────┴──────────────────────┴─────┴─────────┴─────────────────────────────────────────┘`
			So("\n"+content, ShouldEqual, expected)
		})
	})
}
