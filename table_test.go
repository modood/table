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

		Println("\n\nBox Drawing:")
		Output(s)

		Println("\nPure Ascii:")
		OutputA(s)
	})
}

func TestTable(t *testing.T) {
	Convey("Get table content", t, func() {
		Convey("non-slice, should panic", func() {
			So(func() {
				Table(nil)
			}, ShouldPanicWith, "sliceconv: param \"slice\" should be on slice value")

			So(func() {
				AsciiTable(nil)
			}, ShouldPanicWith, "sliceconv: param \"slice\" should be on slice value")
		})

		Convey("slice of int, should panic", func() {
			s := []int{1, 2, 3}

			So(func() {
				Table(s)
			}, ShouldPanicWith, "Table: items of slice should be on struct value")

			So(func() {
				AsciiTable(s)
			}, ShouldPanicWith, "Table: items of slice should be on struct value")
		})

		Convey("special characters should work correctly", func() {
			type S struct {
				Celsius string
			}

			s := []S{{"39.5℃"}, {"37.34℃"}}

			var tb = []struct {
				content  string
				expected string
			}{
				{Table(s), `
┌──────────┐
│ Celsius  │
├──────────┤
│ 39.5℃    │
│ 37.34℃   │
└──────────┘`},
				{AsciiTable(s), `
+----------+
| Celsius  |
+----------+
| 39.5℃    |
| 37.34℃   |
+----------+`},
			}

			for _, tt := range tb {
				So("\n"+tt.content, ShouldEqual, tt.expected)
			}
		})

		Convey("unexported field should be ignored", func() {
			type Field struct {
				Exported   float32
				unexported float32
			}

			s := []Field{{1.2, 2.4}, {4.8, 9.6}}

			var tb = []struct {
				content  string
				expected string
			}{
				{Table(s), `
┌──────────┐
│ Exported │
├──────────┤
│ 1.2      │
│ 4.8      │
└──────────┘`},
				{AsciiTable(s), `
+----------+
| Exported |
+----------+
| 1.2      |
| 4.8      |
+----------+`},
			}

			for _, tt := range tb {
				So("\n"+tt.content, ShouldEqual, tt.expected)
			}
		})

		Convey("Field: string(Chinese)", func() {
			type Poet struct {
				Name    string
				Dynasty string
				Live    string
			}

			s := []Poet{
				{"李白", "唐朝", "701年-762年"},
				{"李清照", "宋朝", "1084年-1155年"},
			}

			var tb = []struct {
				content  string
				expected string
			}{
				{Table(s), `
┌───────────┬─────────┬─────────────────┐
│ Name      │ Dynasty │ Live            │
├───────────┼─────────┼─────────────────┤
│ 李白      │ 唐朝    │ 701年-762年     │
│ 李清照    │ 宋朝    │ 1084年-1155年   │
└───────────┴─────────┴─────────────────┘`},
				{AsciiTable(s), `
+-----------+---------+-----------------+
| Name      | Dynasty | Live            |
+-----------+---------+-----------------+
| 李白      | 唐朝    | 701年-762年     |
| 李清照    | 宋朝    | 1084年-1155年   |
+-----------+---------+-----------------+`},
			}

			for _, tt := range tb {
				So("\n"+tt.content, ShouldEqual, tt.expected)
			}
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

			var tb = []struct {
				content  string
				expected string
			}{
				{Table(s), `
┌────┬──────────────────────┬─────┬─────────┬─────────────────────────────────────────┐
│ ID │ Name                 │ Age │ Deleted │ Created                                 │
├────┼──────────────────────┼─────┼─────────┼─────────────────────────────────────────┤
│ 8  │ Captain Jack Sparrow │ 31  │ false   │ 2017-11-08 23:12:43.249437302 +0000 UTC │
│ 9  │ William Turner       │ 18  │ false   │ 2009-02-09 02:04:25.363779979 +0000 UTC │
│ 10 │ Davy Jones           │ 120 │ true    │ 1965-10-11 07:10:26.106273532 +0000 UTC │
└────┴──────────────────────┴─────┴─────────┴─────────────────────────────────────────┘`},
				{AsciiTable(s), `
+----+----------------------+-----+---------+-----------------------------------------+
| ID | Name                 | Age | Deleted | Created                                 |
+----+----------------------+-----+---------+-----------------------------------------+
| 8  | Captain Jack Sparrow | 31  | false   | 2017-11-08 23:12:43.249437302 +0000 UTC |
| 9  | William Turner       | 18  | false   | 2009-02-09 02:04:25.363779979 +0000 UTC |
| 10 | Davy Jones           | 120 | true    | 1965-10-11 07:10:26.106273532 +0000 UTC |
+----+----------------------+-----+---------+-----------------------------------------+`},
			}

			for _, tt := range tb {
				So("\n"+tt.content, ShouldEqual, tt.expected)
			}
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

			var tb = []struct {
				content  string
				expected string
			}{
				{Table(s), `
┌─────────────────┬────────┬──────────┬───────┬────────────────┬──────────────────────────────┬────────────────────────┐
│ String          │ Int    │ Float    │ Bool  │ Slice          │ Struct                       │ Map                    │
├─────────────────┼────────┼──────────┼───────┼────────────────┼──────────────────────────────┼────────────────────────┤
│ House Stark     │ 100210 │ 3.14159  │ false │ [Arya Bran]    │ {ID:10010 Name:Jon Snow}     │ map[10086:Sansa Stark] │
│ House Lannister │ 2131   │ 1.234234 │ true  │ [Tywin Tyrion] │ {ID:23019 Name:Queen Cersei} │ map[33489:Ser Jaime]   │
└─────────────────┴────────┴──────────┴───────┴────────────────┴──────────────────────────────┴────────────────────────┘`},
				{AsciiTable(s), `
+-----------------+--------+----------+-------+----------------+------------------------------+------------------------+
| String          | Int    | Float    | Bool  | Slice          | Struct                       | Map                    |
+-----------------+--------+----------+-------+----------------+------------------------------+------------------------+
| House Stark     | 100210 | 3.14159  | false | [Arya Bran]    | {ID:10010 Name:Jon Snow}     | map[10086:Sansa Stark] |
| House Lannister | 2131   | 1.234234 | true  | [Tywin Tyrion] | {ID:23019 Name:Queen Cersei} | map[33489:Ser Jaime]   |
+-----------------+--------+----------+-------+----------------+------------------------------+------------------------+`},
			}

			for _, tt := range tb {
				So("\n"+tt.content, ShouldEqual, tt.expected)
			}
		})
	})
}
