table
=====

[![Build Status](https://travis-ci.org/modood/table.png)](https://travis-ci.org/modood/table)
[![Coverage Status](https://coveralls.io/repos/github/modood/table/badge.svg?branch=master)](https://coveralls.io/github/modood/table?branch=master)
[![GoDoc](https://godoc.org/github.com/modood/table?status.svg)](http://godoc.org/github.com/modood/table)

Produces a string that represents slice of structs data in a text table, inspired by gajus/table.

**Features:**

-   No dependency.
-   Cell content aligned.
-   Column width self-adaptation
-   Support type of struct field: int, float, string, bool, slice, struct, map, time.Time and everything.

Installation
------------

```
$ go get github.com/modood/table
```

Quick start
-----------

```go
package main

import (
	"fmt"

	"github.com/modood/table"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}

func main() {
	s := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	t := table.Table(s)

	fmt.Println(t)
}
```

output:
```
┌───────────┬──────────┬──────────────────┐
│ Name      │ Sigil    │ Motto            │
├───────────┼──────────┼──────────────────┤
│ Stark     │ direwolf │ Winter is coming │
│ Targaryen │ dragon   │ Fire and Blood   │
│ Lannister │ lion     │ Hear Me Roar     │
└───────────┴──────────┴──────────────────┘
```

Document
--------

-   `func Output(slice interface{})`

    formats slice of structs data and writes to standard output.

-   `func Table(slice interface{}) string`

    formats slice of structs data and returns the resulting string.

Contributing
------------

1.  Fork it
2.  Create your feature branch (`git checkout -b my-new-feature`)
3.  Commit your changes (`git commit -am 'Add some feature'`)
4.  Push to the branch (`git push origin my-new-feature`)
5.  Create new Pull Request

License
-------

this repo is released under the [MIT License](http://www.opensource.org/licenses/MIT).
