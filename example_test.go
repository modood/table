// Copyright 2017 modood. All rights reserved.
// license that can be found in the LICENSE file.

package table_test

import (
	"github.com/modood/table"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}

func Example() {
	hs := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	// Output to stdout
	table.Output(hs)

	// Or just return table string and then do something
	s := table.Table(hs)
	_ = s
}
