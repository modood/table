// Copyright 2017 modood. All rights reserved.
// license that can be found in the LICENSE file.

package table_test

import (
	"testing"
	"time"

	. "github.com/modood/table"
	. "github.com/smartystreets/goconvey/convey"
)

func TestOutput(t *testing.T) {
	Convey("Output content without panic", t, func() {
		type Repo struct {
			Name   string
			Author string
			URL    string
		}

		s := []Repo{
			{"table", "modood", "https://github.com/modood/table"},
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
			type Field struct {
				Exported   float32
				unexported float32
			}

			s := []Field{{1.2, 2.4}, {4.8, 9.6}}

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

		Convey("Field: string, int, bool, time.Time", func() {
			type User struct {
				ID      string
				Name    string
				Age     int
				Deleted bool
				Created time.Time
			}

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

		Convey("Field: float, slice, struct, map", func() {
			type Mixed struct {
				String string
				Int    int
				Float  float32
				Bool   bool
				Slice  []string
				Struct struct {
					ID   string
					Name string
				}
				Map map[string]string
			}

			s := []Mixed{
				{"House Stark", 100210, 3.14159, false, []string{"Arya", "Bran"}, struct {
					ID   string
					Name string
				}{"10010", "Jon Snow"}, map[string]string{"10086": "Sansa Stark"}},
				{"House Lannister", 2131, 1.234234, true, []string{"Tywin", "Tyrion"}, struct {
					ID   string
					Name string
				}{"23019", "Queen Cersei"}, map[string]string{"33489": "Ser Jaime"}},
			}

			content := Table(s)
			expected := `
┌─────────────────┬────────┬──────────┬───────┬────────────────┬──────────────────────────────┬────────────────────────┐
│ String          │ Int    │ Float    │ Bool  │ Slice          │ Struct                       │ Map                    │
├─────────────────┼────────┼──────────┼───────┼────────────────┼──────────────────────────────┼────────────────────────┤
│ House Stark     │ 100210 │ 3.14159  │ false │ [Arya Bran]    │ {ID:10010 Name:Jon Snow}     │ map[10086:Sansa Stark] │
│ House Lannister │ 2131   │ 1.234234 │ true  │ [Tywin Tyrion] │ {ID:23019 Name:Queen Cersei} │ map[33489:Ser Jaime]   │
└─────────────────┴────────┴──────────┴───────┴────────────────┴──────────────────────────────┴────────────────────────┘`
			So("\n"+content, ShouldEqual, expected)
		})
	})
}
